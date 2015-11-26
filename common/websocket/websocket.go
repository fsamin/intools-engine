package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/gin-gonic/gin"
	"github.com/fsamin/intools-engine/common/logs"
)

var (
	appclient = &AppClient{
		Clients:make(map[*websocket.Conn]*Client),
	}
	wsupgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Client struct {
	Socket			*websocket.Conn
	GroupIds 	[]string
}

type AppClient struct {
	Clients	map[*websocket.Conn]*Client
}

type Message struct{

	Command				string						`json:"Command"`
	Data			map[string]interface{}		`json:"data"`

}

// Registers the connection from to the intools engine
func (appClient *AppClient) Register(conn *websocket.Conn) {

	// Add the client to the connected clients
	logs.Debug.Printf("Connection event from client")

	var client = &Client{
		Socket: conn,
		GroupIds: []string{},
	}
	logs.Debug.Printf("clients before %v \n", appClient.Clients)

	appClient.Clients[conn] = client
	// appClient.Clients = append(appClient.Clients, client)

	logs.Debug.Printf("clients %v \n", appClient.Clients)

	// Send ack
	message := Message{
		Command:	"connected",
		Data:	nil,
	}
	conn.WriteJSON(message)

	// Handles events from client
	Events:
	for {
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			switch err.(type) {
			case *websocket.CloseError:
				logs.Debug.Printf("Websocket %p is deconnected. Removing from clients", conn)
				delete(appclient.Clients, conn)
				logs.Debug.Printf("Clients are now  %v", appClient.Clients)
				break Events
			default:
				logs.Error.Printf("Error while reading json message : %s", err)
				continue Events
			}
		}

		logs.Debug.Printf("Message %v", message)

		switch message.Command {
		case "register-group":
			// Handles group registering for the client
			client.GroupIds = append(client.GroupIds, message.Data["groupId"].(string))
			logs.Debug.Printf("Registered group %s for client %p", message.Data["groupId"], client)
		case "unregister-group":
			// Handles group unregistering for the client
			client.GroupIds = del(client.GroupIds, message.Data["groupId"].(string))
			logs.Debug.Printf("Unregistered group %s for client %p", message.Data["groupId"], client)
		}

		logs.Debug.Printf("Registered groups for client %p are now : %s", client, client.GroupIds)

	}

}

// Broadcasts the value to all client registred to the group
func Notify(groupId string, value *map[string]interface{}) {

	logs.Debug.Printf("Notify all client registred to groupid %s , with value %v", groupId, value)

	logs.Debug.Printf("clients %v \n", appclient.Clients)

	for _,client := range appclient.Clients {

		if contains(client.GroupIds, groupId) {
			data := map[string]interface{}{
			"groupId" : groupId,
			"value" : value,
			}
			message := Message{
				Command:	"connector-value",
				Data:	data,
			}
			logs.Debug.Printf("Notifying client %p with message %s", client, message)
			client.Socket.WriteJSON(message)
		}
	}
}

// Get websocket
func GetWS(c *gin.Context){

	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	// Registering the connection from Intools-front
	appclient.Register(conn)
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func del(slice[]string, item string) []string {
	var r[]string
	for _, str := range slice {
		if str != item {
			r = append(r, str)
		}
	}
	return r
}