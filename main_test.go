package main

import (
	"fmt"
	gnet "net"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"

	"github.com/dedis/drand/core"
	"github.com/dedis/drand/fs"
	"github.com/dedis/drand/key"
	"github.com/dedis/drand/test"
	"github.com/dedis/kyber/pairing/bn256"
	"github.com/dedis/kyber/share"
	"github.com/kabukky/httpscerts"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	cmd := exec.Command("go", "install")
	cmd.Run()
	code := m.Run()
	os.Exit(code)
}

func TestKeyGen(t *testing.T) {
	tmp := path.Join(os.TempDir(), "drand")
	defer os.RemoveAll(tmp)
	cmd := exec.Command("drand", "--folder", tmp, "generate-keypair", "127.0.0.1:8081")
	out, err := cmd.Output()
	require.Nil(t, err)
	fmt.Println(string(out))
	config := core.NewConfig(core.WithConfigFolder(tmp))
	fs := key.NewFileStore(config.ConfigFolder())
	priv, err := fs.LoadKeyPair()
	require.Nil(t, err)
	require.NotNil(t, priv.Public)

	tmp2 := path.Join(os.TempDir(), "drand2")
	defer os.RemoveAll(tmp2)
	cmd = exec.Command("drand", "--folder", tmp2, "generate-keypair")
	out, err = cmd.Output()
	require.Error(t, err)
	fmt.Println(string(out))
	config = core.NewConfig(core.WithConfigFolder(tmp2))
	fs = key.NewFileStore(config.ConfigFolder())
	priv, err = fs.LoadKeyPair()
	require.Error(t, err)
	require.Nil(t, priv)
}

//tests valid commands and then invalid commands
func TestGroup(t *testing.T) {
	n := 5
	tmpPath := path.Join(os.TempDir(), "drand")
	os.Mkdir(tmpPath, 0740)
	defer os.RemoveAll(tmpPath)

	names := make([]string, n, n)
	privs := make([]*key.Pair, n, n)
	for i := 0; i < n; i++ {
		names[i] = path.Join(tmpPath, fmt.Sprintf("drand-%d.public", i))
		privs[i] = key.NewKeyPair("127.0.0.1")
		require.NoError(t, key.Save(names[i], privs[i].Public, false))
		if yes, err := fs.Exists(names[i]); !yes || err != nil {
			t.Fatal(err.Error())
		}
	}

	//test not enough keys
	cmd := exec.Command("drand", "--folder", tmpPath, "group", names[0])
	out, err := cmd.CombinedOutput()
	expectedOut := "group command take at least 3 keys as arguments"
	fmt.Println(string(out))
	require.Error(t, err)

	//test valid creation
	groupPath := path.Join(tmpPath, key.GroupFolderName)
	args := []string{"drand", "--folder", tmpPath, "group"}
	args = append(args, names...)
	cmd = exec.Command(args[0], args[1:]...)
	out, err = cmd.Output()
	expectedOut = "Group file written in /tmp/drand/groups. Distribute it to all the participants to start the DKG."
	fmt.Println(string(out))
	require.True(t, strings.Contains(string(out), expectedOut))
	require.Nil(t, err)
	group := new(key.Group)
	require.NoError(t, key.Load(groupPath, group))
	for i := 0; i < n; i++ {
		require.True(t, group.Contains(privs[i].Public))
	}
	extraName := path.Join(tmpPath, fmt.Sprintf("drand-%d.public", n))
	extraPriv := key.NewKeyPair("127.0.0.1")
	require.NoError(t, key.Save(extraName, extraPriv.Public, false))
	if yes, err := fs.Exists(extraName); !yes || err != nil {
		t.Fatal(err.Error())
	}

	//test valid merge
	cmd = exec.Command("drand", "--folder", tmpPath, "group", "--group", groupPath, extraName)
	out, err = cmd.Output()
	fmt.Println(string(out))
	expectedOut = "Group file updated can be found at /tmp/drand/groups. Run upgrade command to do the resharing."
	require.True(t, strings.Contains(string(out), expectedOut))
	group = new(key.Group)
	require.NoError(t, key.Load(groupPath, group))
	for i := 0; i < n; i++ {
		require.True(t, group.Contains(privs[i].Public))
	}
	require.True(t, group.Contains(extraPriv.Public))

	//test there is already a group file
	cmd = exec.Command(args[0], args[1:]...)
	out, err = cmd.Output()
	fmt.Println(string(out))
	expectedOut = "drand: there already is a group.toml file, please use the flag --group to merge your new keys."
	require.Error(t, err)
	require.True(t, strings.Contains(string(out), expectedOut))

	//test could not load group file
	wrongGroupPath := "not_here"
	cmd = exec.Command("drand", "--folder", tmpPath, "group", "--group", wrongGroupPath, names[0])
	out, err = cmd.CombinedOutput()
	fmt.Println(string(out))
	require.Error(t, err)
}

func TestStartAndStop(t *testing.T) {
	tmpPath := path.Join(os.TempDir(), "drand")
	os.Mkdir(tmpPath, 0740)
	defer os.RemoveAll(tmpPath)
	varEnv := "CRASHCRASH"
	n := 5
	_, group := test.BatchIdentities(n)
	groupPath := path.Join(tmpPath, fmt.Sprintf("group.toml"))
	require.NoError(t, key.Save(groupPath, group, false))

	cmd := exec.Command("drand", "--folder", tmpPath, "start", groupPath, "--tls-disable")
	cmd.Env = append(os.Environ(), varEnv+"=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && e.Success() {
		t.Fatal(err)
	}
	cmd = exec.Command("drand", "-c", tmpPath, "stop")
	cmd.Env = append(os.Environ(), varEnv+"=1")
	err = cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && e.Success() {
		t.Fatal(err)
	}
}

func TestStartBeacon(t *testing.T) {
	tmpPath := path.Join(os.TempDir(), "drand")
	os.Mkdir(tmpPath, 0740)
	defer os.RemoveAll(tmpPath)
	varEnv := "CRASHCRASH"
	n := 5
	_, group := test.BatchIdentities(n)
	groupPath := path.Join(tmpPath, fmt.Sprintf("group.toml"))
	require.NoError(t, key.Save(groupPath, group, false))

	cmd := exec.Command("drand", "--folder", tmpPath, "start", "--tls-disable")
	cmd.Env = append(os.Environ(), varEnv+"=1")
	out, err := cmd.Output()
	fmt.Print(string(out))
	if e, ok := err.(*exec.ExitError); ok && e.Success() {
		t.Fatal(err)
	}
}

func TestStartWithoutGroup(t *testing.T) {
	tmpPath := path.Join(os.TempDir(), "drand")
	os.Mkdir(tmpPath, 0740)
	defer os.RemoveAll(tmpPath)

	cmd := exec.Command("drand", "--folder", tmpPath, "start", "--tls-disable")
	out, err := cmd.Output()
	expectedErr := "The DKG has not been run before, please provide a group file to do the setup."
	output := string(out)
	require.Error(t, err)
	require.True(t, strings.Contains(output, expectedErr))
}

func TestClientTLS(t *testing.T) {
	tmpPath := path.Join(os.TempDir(), "drand")
	os.Mkdir(tmpPath, 0740)
	defer os.RemoveAll(tmpPath)

	pubPath := path.Join(tmpPath, "pub.key")
	certPath := path.Join(tmpPath, "server.pem")
	keyPath := path.Join(tmpPath, "key.pem")

	addr := "127.0.0.1:8082"

	priv := key.NewTLSKeyPair(addr)
	require.NoError(t, key.Save(pubPath, priv.Public, false))

	config := core.NewConfig(core.WithConfigFolder(tmpPath))
	fs := key.NewFileStore(config.ConfigFolder())
	fs.SaveKeyPair(priv)

	if httpscerts.Check(certPath, keyPath) != nil {
		fmt.Println("generating on the fly")
		h, _, _ := gnet.SplitHostPort(priv.Public.Address())
		if err := httpscerts.Generate(certPath, keyPath, h); err != nil {
			panic(err)
		}
	}

	// fake group
	_, group := test.BatchTLSIdentities(5)
	group.Nodes[0] = &key.IndexedPublic{
		Identity: priv.Public,
		Index:    0,
	}
	groupPath := path.Join(tmpPath, fmt.Sprintf("groups/drand_group.toml"))
	fs.SaveGroup(group)

	// fake dkg outuput
	keyStr := "012067064287f0d81a03e575109478287da0183fcd8f3eda18b85042d1c8903ec8160c56eb6d5884d8c519c30bfa3bf5181f42bcd2efdbf4ba42ab0f31d13c97e9552543be1acf9912476b7da129d7c7e427fbafe69ac5b635773f488b8f46f3fc40c673b93a08a20c0e30fd84de8a89adb6fb95eca61ef2fff66527b3be4912de"
	fakeKey, _ := test.StringToPoint(keyStr)
	distKey := &key.DistPublic{Key: fakeKey}
	require.NoError(t, fs.SaveDistPublic(distKey))

	//fake share
	pairing := bn256.NewSuite()
	scalarOne := pairing.G2().Scalar().One()
	s := &share.PriShare{I: 2, V: scalarOne}
	share := &key.Share{Share: s}
	fs.SaveShare(share)

	os.Args = []string{"drand", "--folder", tmpPath, "start", "--tls-cert", certPath, "--tls-key", keyPath, groupPath}
	go main()

	installCmd := exec.Command("go", "install")
	_, err := installCmd.Output()
	require.NoError(t, err)

	cmd := exec.Command("drand", "get", "private", "--tls-cert", certPath, "--nodes", addr, groupPath)
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	require.NoError(t, err)

	/*cmd = exec.Command("drand", "get", "public", "--tls-cert", certPath, "--nodes", addr, groupPath)
	out, err = cmd.CombinedOutput()
	fmt.Println(string(out))
	require.NoError(t, err)*/

	cmd = exec.Command("drand", "get", "cokey", "--tls-cert", certPath, "--nodes", addr)
	out, err = cmd.CombinedOutput()
	fmt.Println(string(out))
	expectedOutput := "ASBnBkKH8NgaA+V1EJR4KH2gGD/Njz7aGLhQQtHIkD7IFgxW621YhNjFGcML+jv1GB9CvNLv2/S6QqsPMdE8l+lVJUO+Gs+ZEkdrfaEp18fkJ/uv5prFtjV3P0iLj0bz/EDGc7k6CKIMDjD9hN6Kia22+5Xsph7y//ZlJ7O+SRLe"
	require.True(t, strings.Contains(string(out), expectedOutput))
	require.NoError(t, err)

	cmd = exec.Command("drand", "show", "share")
	out, err = cmd.CombinedOutput()
	fmt.Println(string(out))
	expectedOutput = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE="
	require.True(t, strings.Contains(string(out), expectedOutput))
	require.NoError(t, err)

	cmd = exec.Command("drand", "show", "public")
	out, err = cmd.CombinedOutput()
	fmt.Println(string(out))
	require.NoError(t, err)

	cmd = exec.Command("drand", "show", "private")
	out, err = cmd.CombinedOutput()
	fmt.Println(string(out))
	require.NoError(t, err)

	cmd = exec.Command("drand", "show", "cokey")
	out, err = cmd.CombinedOutput()
	fmt.Println(string(out))
	expectedOutput = "ASBnBkKH8NgaA+V1EJR4KH2gGD/Njz7aGLhQQtHIkD7IFgxW621YhNjFGcML+jv1GB9CvNLv2/S6QqsPMdE8l+lVJUO+Gs+ZEkdrfaEp18fkJ/uv5prFtjV3P0iLj0bz/EDGc7k6CKIMDjD9hN6Kia22+5Xsph7y//ZlJ7O+SRLe"
	require.True(t, strings.Contains(string(out), expectedOutput))
	require.NoError(t, err)
}
