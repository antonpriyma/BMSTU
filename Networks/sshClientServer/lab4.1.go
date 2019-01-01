package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"

	"os"

	"golang.org/x/crypto/ssh"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <user> <host:port>", os.Args[0])
	}
	i:=3
		client,_, _ := connectToHost(os.Args[1], os.Args[2])
	for true {
		session,err:=client.NewSession()
		if err != nil {
			client.Close()
		}
		if err != nil {
			panic(err)
		}
		var stdoutBuf bytes.Buffer
		session.Stdout = &stdoutBuf
		buffer:=bufio.NewReader(os.Stdin)
		cmd,_:=buffer.ReadString('\n')
		session.Run(cmd)
		i+=1
		fmt.Print(stdoutBuf.String())

	}
		client.Close()
}

func connectToHost(user, host string) (*ssh.Client, *ssh.Session, error) {
	var pass string
	fmt.Print("Password: ")
	fmt.Scanf("%s\n", &pass)

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}