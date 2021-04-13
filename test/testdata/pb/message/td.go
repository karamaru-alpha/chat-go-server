package testdata

import (
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"

	tdCommonString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/common"
	tdMessageString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/message"
)

// Message メッセージのgRPC型テストデータ
var Message = pb.Message{
	Id:     tdCommonString.ULID.Valid,
	RoomId: tdCommonString.ULID.Valid,
	Body:   tdMessageString.Body.Valid,
}
