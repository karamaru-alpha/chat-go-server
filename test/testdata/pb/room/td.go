package testdata

import (
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"

	tdCommonString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/common"
	tdRoomString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/room"
)

// Room トークルームのgRPC型テストデータ
var Room = pb.Room{
	Id:    tdCommonString.ULID.Valid,
	Title: tdRoomString.Title.Valid,
}
