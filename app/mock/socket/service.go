package socket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WebSocketConnection struct {
	*websocket.Conn
}

type Room struct {
	Connections []WebSocketConnection
}

type WsConnectionsHandler struct {
	Rooms map[string]*Room
}

var WsConnections = WsConnectionsHandler{
	make(map[string]*Room),
}

func (wch *WsConnectionsHandler)DeleteRoom(key string) {
	delete(wch.Rooms, key)
}


func (wch *WsConnectionsHandler)BroadcastAll(msg string) {
	for _, room := range wch.Rooms {
		room.Broadcast(msg)
	}
}

func (wc *WebSocketConnection)Broadcast(msg interface{}) {
	err := wc.WriteJSON(msg)
	if err != nil {
		log.Println("Web socket broadcast error.")
		wc.Close()
	}
}


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (wch *WsConnectionsHandler)WsEndpoint(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}

	conn := WebSocketConnection{
		Conn: ws,
	}

	keys, ok := c.Request.URL.Query()["roomId"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Missing roomId")
		return
	}
	roomName := keys[0]
	fmt.Println(roomName)
	fmt.Println("CONNECT")

	if room, ok := wch.Rooms[roomName]; ok {
		fmt.Println(room.Connections)
		fmt.Println("ROOM NALEZEN")

		room.Join(conn)
	} else {
		fmt.Println("ROOM NENALEZEN")
		wch.Rooms[roomName] = &Room{make([]WebSocketConnection, 0)}
		wch.Rooms[roomName].Join(conn)
	}
}

func (r *Room)Broadcast(msg interface{}) {
	fmt.Println("BROADCAST TO POCKET")

	for _, conn := range r.Connections {
		fmt.Println("BROADCAST")
		fmt.Println(msg)
		conn.Broadcast(msg)
	}
}

func (r *Room)Join(conn WebSocketConnection) {
	r.Connections = append(r.Connections, conn)
	log.Println("Client connected to endpoint")
}




