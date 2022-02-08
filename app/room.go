package app

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID                      primitive.ObjectID `json:"_id,omitempty`
	Name                    string             `json:"name"`
	BroadcastChannel        chan *WsPayload
	ClientJoinChannel       chan *WsPayload
	ClientDisconnectChannel chan *WsPayload
	OperationChannel        chan *WsPayload
	ActiveConnections       map[WebSocketConnection]string
}

type RoomPersistence struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" sql:"room_id"`
	Name        string             `bson:"name" sql:"name"`
	Subscribers []string           `bson:"subscribers" sql:"subscribers"`
	Active      bool               `bson:"active" sql:"active"`
	Private     bool               `bson:"private" sql:"private"`
	Owner       string             `bson:"owner" sql:"owner"`
}

type WsServer struct {
	ID                      primitive.ObjectID `json:"_id,omitempty`
	Name                    string             `json:"name"`
	BroadcastChannel        chan *WsPayload
	ClientJoinChannel       chan *WsPayload
	ClientDisconnectChannel chan *WsPayload
	OperationChannel        chan *WsPayload
	ActiveConnections       map[WebSocketConnection]string
	Rooms                   map[*Room]bool
	Organization            string
}

type WsServerPersistence struct {
	ID           primitive.ObjectID   `json:"_id,omitempty`
	Name         string               `json:"name"`
	Active       bool                 `bson:"active" sql:"active"`
	Private      bool                 `bson:"private" sql:"private"`
	Organization string               `bson:"organization" sql:"organization"`
	Rooms        []primitive.ObjectID `bson:"rooms" sql:"room_id"`
}

// wsServer.reactivateRooms
// wsServer.createRoom
// wsServer.deleteRoom

// room.loadHistory
