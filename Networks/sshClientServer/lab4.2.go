package main

import (
	"bytes"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"log"
	"os/exec"
	"strings"

	"github.com/gliderlabs/ssh"
)

var (
	password = flag.String("p", "", "Password")
)

func main() {

	ssh.Handle(func(s ssh.Session) {
		fmt.Println(s)
		var cmd *exec.Cmd
		if len(s.Command()) == 0 {
			term := terminal.NewTerminal(s, "> ")
			for {

				line, err := term.ReadLine()
				if err != nil && err != io.EOF {
					log.Println(err.Error() + line)
					break
				}
				lines := strings.Split(line, " ")
				if len(lines) == 1 {
					cmd = exec.Command(lines[0])
				} else {
					cmd = exec.Command(lines[0], lines[1:]...)
				}

				var out bytes.Buffer
				cmd.Stdout = &out
				err = cmd.Run()
				if err != nil {
					term.Write(append([]byte(err.Error()), '\n'))
				}
				log.Println("response: ", cmd)
				if line != "" {
					log.Println(line)
				}
				if out.String() != "" {
					term.Write(append([]byte(out.String()), '\n'))
				}
			}
		} else {
			if len(s.Command()) > 1 {
				log.Println(s.Command())
				cmd = exec.Command(s.Command()[0], s.Command()[1:]...)
			} else if len(s.Command()) == 1 {
				cmd = exec.Command(s.Command()[0])
			}

			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err == nil {
				io.WriteString(s, fmt.Sprintf("Hello %s\n", out.String()))
			} else {
				io.WriteString(s, "cmd.Run() failed with "+err.Error()+"\n")
			}
		}

	})
	flag.Parse()
	log.Println(*password)
	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil, ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
		return pass == *password
	})))
}
