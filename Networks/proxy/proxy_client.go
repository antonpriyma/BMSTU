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
	"sync"
)



// chat server command /command [username] body contents
var chatServerResponseRegex, _ = regexp.Compile(`^\/([^\s]*)\s?(?:\[([^\]]*)\])?\s*(.*)$`)
var diskResponseRegex, _ = regexp.Compile(`\S*\b%?`)
var serverResponseRegex, _ = regexp.Compile(`\S*\b`)

var sshHosts=[]string{"lab.posevin.com:22","lab2.posevin.com:22"}
var sshNames=[]string{"lab_52_2018","lab2_52_2018"}
var sshPasswords=[]string{"fdfrfdfgfdf","fdfrfdfgfdf"}
var serverInfo ServerInfo


type Command struct {
	//Структура команды, куда также входит и message
	Command,Username,Body string
}

type DiskInfo struct {
	Size,Used,Avail,UsedPros string
}

type ServerInfo struct {
	sshHost, sshName,sshPassword,userName string
}

func main()  {
	username:=createConfig()
	conn,err:=net.Dial("tcp","192.168.43.67:9000")
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
		username:="go"
		log.Info("Username:",username)
		return username
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
				sendCommand("goClient", username, conn)

			case "connect":
				fmt.Println( Command.Username,"has connected to lobby")
				parseServer(Command.Body)
				log.Info(serverInfo.sshHost)
				log.Info(serverInfo.sshName)
				log.Info(serverInfo.sshPassword)

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
			case "getDiskInfo":
				info:=executeDiskInfoCommand()
				sendCommand("postDiskInfo",info,conn)
			case "exe":
				c:=make(chan  string)
				log.Info(Command.Body)
				executeCommand(Command.Body)
				sendCommandToServers(Command.Body,c)
				reply:=<-c
				log.Info("reply ",reply)
				sendCommand("postExe",reply,conn)
			case "filesExe":
				c:=make(chan  string)
				log.Info(Command.Body)
				executeCommand(Command.Body)
				sendCommandToServers(Command.Body,c)
				reply:=<-c
				log.Info("reply ",reply)
				sendCommand("postFilesExe",reply,conn)


				


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
				case "test":






				default:
					fmt.Printf("Unknown command \"%s\"\n", command.Command)
				}
			}
		}

		}

	}


func sendCommand(command string, body string, conn net.Conn) {
	message := fmt.Sprintf("/%v %v", command, body)
	log.Info("Message:",strings.Replace(message,"\n","$",-1))
	log.Info("Message byte:",string([]byte(message)))
	//conn.Write([]byte(message))
	fmt.Fprintln(conn,strings.Replace(message,"\n","$",-1))
}



func parseInput(message string) Command {
	res:=strings.Split(message," ")
	if res[0]=="/test" {
		//testCommand:=""
		//for _,s:=range res[1:len(res)-2]{
		//	testCommand+= s+" "
		//}
		//testCommand+=res[len(res)-1]
		return Command{
			Command:res[0][1:],
			Body:strings.Replace(res[1],"\n","",1),
		}
	}else
	if len(res)==2 && strings.Contains(message,"/"){
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

func parseServer(message string){
	res:=serverResponseRegex.FindAllString(message,-1)
	log.Info("res[0]",res[0])
	serverInfo.sshHost=res[0]
	serverInfo.sshName=res[1]
	serverInfo.sshPassword=res[2]
	serverInfo.userName="android"



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

func execCmd(client *ssh.Client, session *ssh.Session,cmd string) (string){
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	log.Info("point 2")
	session.Run(cmd)

	log.Info(stdoutBuf.String())
	log.Info("point 2")
	client.Close()
	return stdoutBuf.String()
}

func executeDiskInfoCommand() (string){

	reply :=make(chan string)
	sendCommandToServers("df -h", reply)

	reply1 :=<-reply
	reply1=strings.Split(reply1,"\n")[1]
	res := diskResponseRegex.FindAllString(reply1, -1)
	log.Info(reply1)
	log.Info("reply1",res[1])
	disk:=fmt.Sprintf("%v %v %v %v %v",res[0],res[1],res[2],res[3],res[4])

	//info1:=DiskInfo{
	//	Size:res[1],
	//	Used:res[2],
	//	Avail:res[3],
	//	UsedPros:res[4],
	//}
	//log.Info(info1.Size)
	//log.Info(info1.Used)
	//log.Info(info1.Avail)
	//log.Info(info1.UsedPros)
	//res = diskResponseRegex.FindAllString(reply2, -1)
	//info2:=DiskInfo{
	//	Size:res[1],
	//	Used:res[2],
	//	Avail:res[3],
	//	UsedPros:res[4],
	//}
	return disk
}

func executeCommand(command string){
	switch command {
		case "diskInfo":



	}
}

func sendCommandToServers(command string,c chan string){
	var wg sync.WaitGroup
	//replies:=make([]string,2)
	wg.Add(1)
	//for i,hostname:=range sshHosts{
	//	sshConfig := &ssh.ClientConfig{
	//		User: sshNames[i],
	//		Auth: []ssh.AuthMethod{ssh.Password(sshPasswords[i])},
	//	}
	//	go func(hostname string){
	//		defer wg.Done()
	//		log.Info("point")
	//		client,session,err:=connectToHosts(sshConfig,hostname)
	//		if err != nil {
	//			return
	//		}
	//		c<-strings.Split(execCmd(client,session,command),"\n")[1]
	//	}(hostname)
	//}
	sshConfig := &ssh.ClientConfig{
		User: serverInfo.sshName,
		Auth: []ssh.AuthMethod{ssh.Password(serverInfo.sshPassword)},
	}
	go func(hostname string){
		defer wg.Done()
		log.Info("point")
		client,session,err:=connectToHosts(sshConfig,hostname)
		if err != nil {
			return
		}
		c<-execCmd(client,session,command)
	}(serverInfo.sshHost)





}






