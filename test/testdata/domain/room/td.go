package testdata

import (
	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

// Room トークルームエンティティにまつわるテストデータ
var Room = struct {
	Entity domainModel.Room
	ID     domainModel.ID
	Title  domainModel.Title
}{
	Entity: genEntity(),
	ID:     genID(),
	Title:  genTitle(),
}

func genEntity() domainModel.Room {
	return domainModel.Room{
		ID:    genID(),
		Title: genTitle(),
	}
}

func genID() domainModel.ID {
	return domainModel.ID(tdULID.Room.ID)
}

func genTitle() domainModel.Title {
	return domainModel.Title(tdString.Room.Title.Valid)
}
