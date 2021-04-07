package testdata

import (
	"github.com/karamaru-alpha/chat-go-server/domain/model/message"
	"github.com/karamaru-alpha/chat-go-server/domain/model/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

// Message メッセージエンティティにまつわるテストデータ
var Message = struct {
	Entity message.Message
	ID     message.ID
	RoomID room.ID
	Body   message.Body
}{
	Entity: genEntity(),
	ID:     genID(),
	RoomID: genRoomID(),
	Body:   genBody(),
}

func genEntity() message.Message {
	return message.Message{
		ID:     genID(),
		RoomID: genRoomID(),
		Body:   genBody(),
	}
}

func genID() message.ID {
	return message.ID(tdULID.Message.ID)
}

func genRoomID() room.ID {
	return room.ID(tdULID.Room.ID)
}

func genBody() message.Body {
	return message.Body(tdString.Message.Body.Valid)
}
