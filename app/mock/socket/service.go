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
	BroadcastRooms map[string]*Room
	MeteringRooms  map[string]*Room
}

var WsConnections = WsConnectionsHandler{
	make(map[string]*Room),
	make(map[string]*Room),
}

func (wch *WsConnectionsHandler) DeleteRoom(roomType string, key string) {
	if roomType == "Broadcast" {
		delete(wch.BroadcastRooms, key)
	} else {
		delete(wch.MeteringRooms, key)
	}
}

func (wch *WsConnectionsHandler) BroadcastAll(msg string) {
	for _, room := range wch.BroadcastRooms {
		room.Broadcast(msg)
	}
	for _, room := range wch.MeteringRooms {
		room.Broadcast(msg)
	}
}

func (wc *WebSocketConnection) Broadcast(msg interface{}) {
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

func (wch *WsConnectionsHandler) WsEndpoint(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}

	conn := WebSocketConnection{
		Conn: ws,
	}

	// todo parse roomType
	roomType := c.Request.URL.Query().Get("roomType")
	roomId := c.Request.URL.Query().Get("roomId")

	fmt.Println(roomId)
	fmt.Println("CONNECT")

	var room *Room
	if roomType == "broadcast" {
		if r, ok := wch.BroadcastRooms[roomId]; ok {
			room = r
		}
	} else if roomType == "metering" {
		if r, ok := wch.MeteringRooms[roomId]; ok {
			fmt.Printf("room %s in %s", roomId, roomType)
			room = r
		}
	} else {
		fmt.Println("unknown roomType")
		return
	}

	if room != nil {
		fmt.Printf("room %s in %s", roomId, roomType)
		room.Join(conn)
		fmt.Printf("socket connected to room %s in %s", roomId, roomType)

	} else {
		fmt.Printf("room not found, creating room %s in %s", roomId, roomType)
		wch.BroadcastRooms[roomId] = &Room{make([]WebSocketConnection, 0)}
		wch.BroadcastRooms[roomId].Join(conn)
		fmt.Printf("socket connected to new room %s in %s", roomId, roomType)
	}
}

func (r *Room) Broadcast(msg interface{}) {
	fmt.Println("BROADCAST TO POCKET")

	for _, conn := range r.Connections {
		fmt.Println("BROADCAST")
		fmt.Println(msg)
		conn.Broadcast(msg)
	}
}

func (r *Room) Join(conn WebSocketConnection) {
	r.Connections = append(r.Connections, conn)
	log.Println("Client connected to endpoint")
}
