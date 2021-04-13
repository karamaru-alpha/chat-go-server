package testdata

import (
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

var Message = pb.Message{
	Id:     tdString.Message.ID.Valid,
	RoomId: tdString.Room.ID.Valid,
	Body:   tdString.Message.Body.Valid,
}
