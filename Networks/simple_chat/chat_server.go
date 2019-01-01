package main

import (
	"./util"
	"bufio"
	"github.com/mgutz/logxi/v1"
	"net"
	"regexp"
	"strings"
)


const LOBBY = "lobby"
func main() {

	listener, _ := net.Listen("tcp", "192.168.1.40:9000")
	log.Info("Chat server started on port :2222\n")

	for {
		conn, _ := listener.Accept()
		log.Info("Connection accepted")

		client:=util.Client{Connection:conn,Room:LOBBY}
		log.Info("Client:",client.Connection.RemoteAddr())
		client.Register()
		channel := make(chan string)
		go waitForInput(channel, &client)
		go handleInput(channel, &client)
		util.SendClientMessage("ready",":2222", &client, true)

	}
}

func waitForInput(out chan string, client *util.Client) {
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


func handleInput(in <-chan string, client *util.Client) {

	for {
		message := <- in
		if (message != "") {
			message = strings.TrimSpace(message)
			action, body := getAction(message)
			log.Info("Action:",action)
			log.Info("Body:",body)

			if (action != "") {
				switch action {


				case "message":
					log.Info("username:",client.Username)
					util.SendClientMessage("message", body, client, false)


				case "user":
					client.Username = body
					util.SendClientMessage("connect", "", client, false)


				case "disconnect":
					client.Close()


				case "ignore":
					client.Ignore(body)
					util.SendClientMessage("ignoring", body, client, false)


				case "enter":
					if (body != "") {
						client.Room = body
						util.SendClientMessage("enter", body, client, false)
					}

				
				case "leave":
					if (client.Room != LOBBY) {
						util.SendClientMessage("leave", client.Room, client, false)
						client.Room = LOBBY
					}

				default:
					util.SendClientMessage("unrecognized", action, client, true)
				}
			}
		}
	}
}

func getAction(message string) (string, string) {
	actionRegex, _ := regexp.Compile(`^\/([^\s]*)\s*(.*)$`)
	res := actionRegex.FindAllStringSubmatch(message, -1)
	if (len(res) == 1) {
		return res[0][1], res[0][2]
	}
	return "", ""
}