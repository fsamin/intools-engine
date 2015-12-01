package websocket

import (
	"fmt"
	"github.com/fsamin/intools-engine/common/logs"
	"github.com/fsamin/intools-engine/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	appclient = &AppClient{
		Clients: make(map[*websocket.Conn]*Client),
	}
	wsupgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Client struct {
	Socket   *websocket.Conn
	GroupIds []string
}

type AppClient struct {
	Clients map[*websocket.Conn]*Client
}

type Message struct {
	Key  string                 `json:"key"`
	Data map[string]interface{} `json:"data"`
}

// Registers the connection from to the intools engine
func (appClient *AppClient) Register(conn *websocket.Conn) {
	// Add the client to the connected clients
	logs.Debug.Printf("Connection event from client")

	var client = &Client{
		Socket:   conn,
		GroupIds: []string{},
	}
	logs.Debug.Printf("clients before %v", appClient.Clients)

	appClient.Clients[conn] = client
	// appClient.Clients = append(appClient.Clients, client)

	logs.Debug.Printf("clients %v", appClient.Clients)

	// Send ack
	message := Message{
		Key:  "connected",
		Data: nil,
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

		switch message.Key {
		case "register-group":
			// Handles group registering for the client
			client.GroupIds = append(client.GroupIds, message.Data["groupId"].(string))
			logs.Debug.Printf("Registered group %s for client %p", message.Data["groupId"], client)
		case "unregister-group":
			// Handles group unregistering for the client
			i := utils.IndexOf(client.GroupIds, message.Data["groupId"].(string))
			if i != -1 {
				client.GroupIds = append(client.GroupIds[:i], client.GroupIds[i+1:]...)
				logs.Debug.Printf("Unregistered group %s for client %p", message.Data["groupId"], client)
			}
		}
		logs.Debug.Printf("Registered groups for client %p are now : %s", client, client.GroupIds)
	}
}

// Broadcasts the value to all client registred to the group
func Notify(groupId string, connectorId string, value *map[string]interface{}) {
	logs.Debug.Printf("Notify all client registred to groupid %s , with value %v", groupId, value)

	logs.Debug.Printf("clients %v \n", appclient.Clients)

	for _, client := range appclient.Clients {

		if utils.Contains(client.GroupIds, groupId) {
			data := map[string]interface{}{
				"connectorId": connectorId,
				"value":       value,
			}
			message := Message{
				Key:  "connector-value",
				Data: data,
			}
			logs.Debug.Printf("Notifying client %p with message %s", client, message)
			client.Socket.WriteJSON(message)
		}
	}
}

// Get websocket
func GetWS(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	// Registering the connection from Intools-front
	appclient.Register(conn)
}
