package util

import (
	"fmt"
	"github.com/mgutz/logxi/v1"
	"net"
)

type ServerInfo struct {

	Connection net.Conn


	UserName string
	ServerName string


	Host string

	Port string

	Password string


}

type DiskInfo struct {
	Size,Used,UsedPros string
}

var clients []*ServerInfo

func contains(c *ServerInfo) bool {
	for _, a := range clients {
		if a.Connection == c.Connection {
			return true
		}
	}
	return false
}

func (client *ServerInfo) Register() bool {
	if (!contains(client)) {
		clients = append(clients, client)
		log.Info("register client",client)
		return true
	}
	return false
}

func (client *ServerInfo) Close()  {
	client.Connection.Close()
	clients=removeEntry(client, clients)
}


func removeEntry(client *ServerInfo, arr []*ServerInfo) []*ServerInfo{
	rtn := arr
	index := -1
	for i, value := range arr {
		if value == client {
			index = i
			break
		}
	}

	if index >= 0 {
		rtn = make([]*ServerInfo, len(arr)-1)
		copy(rtn, arr[:index])
		copy(rtn[index:], arr[index+1:])
	}

	return rtn
}

func SendClientMessage(messageType string, message string, client *ServerInfo, forGoClient bool) {

	if forGoClient {
		payload := fmt.Sprintf("/%v [%v] %v", messageType, "android", message)
		//fmt.Fprintln(client.Connection, message)
		log.Info("message util:",payload)
		for _,_client := range clients {
			if _client.UserName=="go"{
				fmt.Fprintln(_client.Connection,payload)
				break
			}
		}

	} else  {

		payload := fmt.Sprintf("/%v [%v] %v", messageType, client.UserName, message)


		for _, _client := range clients {
			if _client.UserName!="go"{
				log.Info("Payload:",payload)
				fmt.Fprintln(_client.Connection, payload)
			}

		}
	}
}

//func (client *Client) Ignore(username string) {
//	client.ignoring = append(client.ignoring, username)
//}

//func (client *Client) IsIgnoring(username string) bool {
//	for _, value := range client.ignoring {
//		if value == username {
//			return true
//		}
//	}
//	return false
//}
