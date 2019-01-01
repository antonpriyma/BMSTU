package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"golang.org/x/crypto/ssh"
)

var (
	user = flag.String("u", "", "User name")
	pk   = flag.String("pk", defaultKeyPath(), "Private key file")
	host = flag.String("h", "", "Host")
	port = flag.Int("p", 22, "Port")
)

func defaultKeyPath() string {
	home := os.Getenv("HOME")
	if len(home) > 0 {
		return path.Join(home, ".ssh/id_rsa")
	}
	return ""
}

func main() {
	flag.Parse()

	key, err := ioutil.ReadFile(*pk)
	if err != nil {
		panic(err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: *user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}

	addr := fmt.Sprintf("%s:%d", *host, *port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic(err)
	}

	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	b, err := session.CombinedOutput("uname -a")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}