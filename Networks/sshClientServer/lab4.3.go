package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"github.com/mgutz/logxi/v1"
	"golang.org/x/crypto/ssh"
)

func main() {

	if len(os.Args) < 4 {
		log.Fatal("Usage: %s <user> <command> <hosts:port>  ", os.Args[0])
	}

	var pass string
	fmt.Print("Password: ")
	fmt.Scanf("%s\n", &pass)
	hosts:=os.Args[3:]
	sshConfig := &ssh.ClientConfig{
		User: os.Args[1],
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}

	var wg sync.WaitGroup
	wg.Add(len(os.Args)-3)
	for _,hostname:=range hosts{
		go func(hostname string){
			defer wg.Done()
			log.Info("point")
			client,session,err:=connectToHosts(sshConfig,hostname)
			if err != nil {
				return
			}
			execCmd(client,session,os.Args[2])
		}(hostname)
	}
	wg.Wait()
}

func connectToHosts(config *ssh.ClientConfig, host string) (*ssh.Client, *ssh.Session, error) {

	config.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Info(host," ",err)
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		log.Info(host)
		client.Close()
		return nil, nil, err
	}
	log.Info("point 1")
	return client, session, nil
}

func execCmd(client *ssh.Client, session *ssh.Session,cmd string){
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(cmd)

	fmt.Print(stdoutBuf.String())
	log.Info("point 2")
	client.Close()
}