package testdata

import (
	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

// Message メッセージエンティティにまつわるテストデータ
var Message = struct {
	Entity messageDomain.Message
	ID     messageDomain.ID
	RoomID roomDomain.ID
	Body   messageDomain.Body
}{
	Entity: genEntity(),
	ID:     genID(),
	RoomID: genRoomID(),
	Body:   genBody(),
}

func genEntity() messageDomain.Message {
	return messageDomain.Message{
		ID:     genID(),
		RoomID: genRoomID(),
		Body:   genBody(),
	}
}

func genID() messageDomain.ID {
	return messageDomain.ID(tdULID.Message.ID)
}

func genRoomID() roomDomain.ID {
	return roomDomain.ID(tdULID.Room.ID)
}

func genBody() messageDomain.Body {
	return messageDomain.Body(tdString.Message.Body.Valid)
}
