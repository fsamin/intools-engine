package websocket

import (
	"github.com/fsamin/intools-engine/common/logs"
	"github.com/fsamin/intools-engine/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	defaultChannelLength = 10
)

var (
	appclient = &AppClient{
		Clients: make(map[*websocket.Conn]*Client),
	}
	wsupgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ConnectorBuffer chan *LightConnector
)

type LightConnector struct {
	GroupId     string
	ConnectorId string
	Value       *map[string]interface{}
}

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

func InitChannel(length int64) {
	if length <= 0 {
		length = defaultChannelLength
	}
	ConnectorBuffer = make(chan *LightConnector, length)
	logs.Trace.Printf("Initializing websocket buffered channel with a size of %+v", length)
	go func() {
		for {
			notify(<-ConnectorBuffer)
		}
	}()
}

// Get websocket
func GetWS(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logs.Error.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	// Registering the connection from Intools back-office
	err = appclient.Register(conn)
	if err != nil {
		switch err.(type) {
		case *websocket.CloseError:
			logs.Trace.Printf("Communication with client has been interrupted : websocket closed")
			return
		default:
			logs.Error.Printf("Error while registering client, closing websocket : %s", err)
			conn.Close()
		}
	}
}

// Registers the connection between a client (one intools backoffice instance) and server (this intools engine instance)
func (appClient *AppClient) Register(conn *websocket.Conn) error {

	client := createClient(conn)
	appclient.bindClient(conn, &client)
	err := sendAck(conn)
	if err != nil {
		logs.Error.Printf("Can't send ack to the client : %s", err)
		return err
	}

	logs.Trace.Printf("Client %v registered", client)

	err = appclient.handleEvents(conn, &client)
	if err != nil {
		return err
	}
	return nil
}

// Broadcasts the value to all client registered to the group
func notify(lConnector *LightConnector) {
	logs.Trace.Printf("Notifying all client registered to groupid %s", lConnector.GroupId)
	logs.Debug.Printf("Value to send : %v , Clients to notify %v", lConnector.Value, appclient.Clients)

	for _, client := range appclient.Clients {
		logs.Debug.Printf("Clients groupids %s", client.GroupIds)
		if utils.Contains(client.GroupIds, lConnector.GroupId) {
			message := createConnectorValueMessage(lConnector.ConnectorId, lConnector.Value)
			logs.Debug.Printf("Notifying client %p with message %s", client, message)
			err := client.Socket.WriteJSON(message)
			if err != nil {
				logs.Warning.Printf("Can't send connector value of id %s, to client %p: %s", lConnector.ConnectorId, client, err)
			}
		}
	}
}

// Creates message, structured as value send to the client
func createConnectorValueMessage(connectorId string, value *map[string]interface{}) Message {
	data := map[string]interface{}{
		"connectorId": connectorId,
		"value":       value,
	}
	message := Message{
		Key:  "connector-value",
		Data: data,
	}
	return message
}

// Create a simple client
func createClient(conn *websocket.Conn) Client {
	logs.Debug.Printf("Connection event from client")
	var client = &Client{
		Socket:   conn,
		GroupIds: []string{},
	}
	return *client
}

// Add the client to the connected clients
func (appClient *AppClient) bindClient(conn *websocket.Conn, c *Client) {
	logs.Debug.Printf("clients before %v", appClient.Clients)
	appClient.Clients[conn] = c
	logs.Debug.Printf("clients %v", appClient.Clients)
}

// Send ack to the client
func sendAck(conn *websocket.Conn) error {
	message := Message{
		Key:  "connected",
		Data: nil,
	}
	err := conn.WriteJSON(message)
	if err != nil {
		return err
	}
	return nil
}

// Handle events from client
func (appClient *AppClient) handleEvents(conn *websocket.Conn, client *Client) error {
Events:
	for {
		// Read message
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			switch err.(type) {
			case *websocket.CloseError:
				logs.Debug.Printf("Websocket %p is deconnected. Removing from clients", conn)
				delete(appclient.Clients, conn)
				logs.Debug.Printf("Clients are now  %v", appClient.Clients)
				return err
			default:
				logs.Warning.Printf("Error while reading json message : %s", err)
				continue Events
			}
		}

		logs.Debug.Printf("Message %v", message)

		// Check message structure
		id, ok := message.Data["groupId"]
		if !ok {
			logs.Warning.Printf("Can't register or unregister group because groupId does not exist in message")
			continue Events
		}
		groupId, ok := id.(string)
		if !ok {
			logs.Warning.Printf("Can't register or unregister group because groupId is not string : %s", groupId)
			continue Events
		}

		// Handles types of messages
		switch message.Key {
		case "register-group":
			// Handles group registering for the client

			client.GroupIds = append(client.GroupIds, groupId)
			logs.Trace.Printf("Registered group %s for client %p", groupId, client)
		case "unregister-group":
			// Handles group unregistering for the client
			i, ok := utils.IndexOf(client.GroupIds, groupId)
			if ok {
				client.GroupIds = append(client.GroupIds[:i], client.GroupIds[i+1:]...)
				logs.Trace.Printf("Unregistered group %s for client %p", groupId, client)
			}
		}
		logs.Debug.Printf("Registered groups for client %p are now : %s", client, client.GroupIds)
	}
}
