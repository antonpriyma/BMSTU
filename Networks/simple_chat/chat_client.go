package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/mgutz/logxi/v1"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
	"regexp"
	"strings"
)



// chat server command /command [username] body contents
var chatServerResponseRegex, _ = regexp.Compile(`^\/([^\s]*)\s?(?:\[([^\]]*)\])?\s*(.*)$`)

type Command struct {
	//Структура команды, куда также входит и message
	Command,Username,Body string
}

func main()  {
	username:=createConfig()
	conn,err:=net.Dial("tcp","192.168.1.40:9000")
	if err!=nil {
		log.Info(err.Error())
		return
	}else {
		log.Info("Connected to tcp ok")
	}
	defer conn.Close()
	log.Info("Point in main 0")
	go watchForConnectionInput(username, conn)
	log.Info("Point in main 1")
	for true{
		log.Info("Point in main 2")
		watchForConsoleInput(conn)
	}


}
func createConfig() (string) {
	if len(os.Args)>=2 {
		username:=os.Args[1]
		log.Info("Username:",username)
		return username
	}else {
		log.Info("Name required")
		os.Exit(0)
		return ""
	}
}

func watchForConnectionInput(username string,  conn net.Conn) {
	reader := bufio.NewReader(conn)

	for true {
		message, _ := reader.ReadString('\n')
		log.Info("message connectioninput",message)
		message = strings.TrimSpace(message)
		if message != "" {
			Command := parseCommand(message)
			log.Info("Command",Command)
			switch Command.Command {

			case "ready":
				sendCommand("user", username, conn)

			case "connect":
				fmt.Println( Command.Username,"has connected to lobby")

			case "disconnect":
				fmt.Printf( Command.Username)

			case "enter":
				fmt.Println( Command.Username,"enter",Command.Body)

			case "leave":
				fmt.Println( Command.Username, "leave",Command.Body)

			case "message":
				log.Info(Command.Username,Command.Body)
				log.Info(username)
				if Command.Username != username {
					fmt.Println(Command.Username, Command.Body)
				}


			case "ignoring":
				fmt.Printf( Command.Body)
			}
		}
	}
}


func watchForConsoleInput(conn net.Conn)  {
	reader := bufio.NewReader(os.Stdin)
	for true{
		message,_:=reader.ReadString('\n')
		log.Info("Message consoleinput",message)
		if message!="" {
			command:=parseInput(message)
			log.Info("Command:",command)
			if command.Command == "" {
				sendCommand("message", message, conn)
			} else {
				switch command.Command {

				case "enter":
					sendCommand("enter", command.Body, conn)

				case "ignore":
					sendCommand("ignore", command.Body, conn)

				case "leave":
					sendCommand("leave", "", conn)

				case "disconnect":
					sendCommand("disconnect", "", conn)

				default:
					fmt.Printf("Unknown command \"%s\"\n", command.Command)
				}
			}
		}

		}

	}


func sendCommand(command string, body string, conn net.Conn) {
	message := fmt.Sprintf("/%v %v\n", command, body)
	log.Info("Message:",message)
	conn.Write([]byte(message))
}



func parseInput(message string) Command {
	res:=strings.Split(message," ")
	if len(res)==2 && strings.Contains(message,"/"){
		//var body string
		//for _,s:=range res{
		//
		//}
		return Command{
			Command:res[0][1:],
			Body:res[1][:len(res[1])-1],
		}
	}else if strings.Contains(message,"/") {
		return Command{
			Command:message[1:len(message)-1],
		}
	}
	return Command{}
}

func parseCommand(message string) Command {
	res := chatServerResponseRegex.FindAllStringSubmatch(message, -1)
	if len(res) == 1 {
		return Command {
			Command: res[0][1],
			Username: res[0][2],
			Body: res[0][3],
		}
	} else {
		return Command{}
	}
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






