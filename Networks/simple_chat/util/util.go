package util

import (
	"fmt"
	"github.com/mgutz/logxi/v1"
	"net"
)

type Client struct {

	Connection net.Conn

	Username string


	Room string

	ignoring []string

	Password string


}

var clients []*Client

func (client *Client) Register() {
	clients = append(clients, client)
}

func (client *Client) Close()  {
	client.Connection.Close()
	clients=removeEntry(client, clients)
}


func removeEntry(client *Client, arr []*Client) []*Client {
	rtn := arr
	index := -1
	for i, value := range arr {
		if value == client {
			index = i
			break
		}
	}

	if index >= 0 {
		rtn = make([]*Client, len(arr)-1)
		copy(rtn, arr[:index])
		copy(rtn[index:], arr[index+1:])
	}

	return rtn
}

func SendClientMessage(messageType string, message string, client *Client, thisClientOnly bool) {

	if thisClientOnly {
		message = fmt.Sprintf("/%v", messageType)
		fmt.Fprintln(client.Connection, message)
		log.Info("message util:",message)

	} else if client.Username != "" {

		payload := fmt.Sprintf("/%v [%v] %v", messageType, client.Username, message)


		for _, _client := range clients {

			if (thisClientOnly && _client.Username == client.Username) ||
							(!thisClientOnly && _client.Username != "") {


				if messageType == "message" && client.Room != _client.Room || _client.IsIgnoring(client.Username) {
					continue
				}

				log.Info("Payload:",payload)
				fmt.Fprintln(_client.Connection, payload)
			}
		}
	}
}

func (client *Client) Ignore(username string) {
	client.ignoring = append(client.ignoring, username)
}

func (client *Client) IsIgnoring(username string) bool {
	for _, value := range client.ignoring {
		if value == username {
			return true
		}
	}
	return false
}
