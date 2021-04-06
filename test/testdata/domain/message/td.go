package testdata

import (
	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

// Message メッセージエンティティにまつわるテストデータ
var Message = struct {
	Entity domainModel.Message
	ID     domainModel.ID
	RoomID domainModel.RoomID
	Body   domainModel.Body
}{
	Entity: genEntity(),
	ID:     genID(),
	RoomID: genRoomID(),
	Body:   genBody(),
}

func genEntity() domainModel.Message {
	return domainModel.Message{
		ID:     genID(),
		RoomID: genRoomID(),
		Body:   genBody(),
	}
}

func genID() domainModel.ID {
	return domainModel.ID(tdULID.Message.ID)
}

func genRoomID() domainModel.RoomID {
	return domainModel.RoomID(tdULID.Room.ID)
}

func genBody() domainModel.Body {
	return domainModel.Body(tdString.Message.Body.Valid)
}
