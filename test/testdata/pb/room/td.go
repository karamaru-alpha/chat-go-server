package testdata

import (
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

var Room = pb.Room{
	Id:    tdString.Room.ID.Valid,
	Title: tdString.Room.Title.Valid,
}
