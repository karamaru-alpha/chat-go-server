package testdata

import (
	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

// Room トークルームエンティティにまつわるテストデータ
var Room = struct {
	Entity domain.Room
	ID     domain.ID
	Title  domain.Title
}{
	Entity: genEntity(),
	ID:     genID(),
	Title:  genTitle(),
}

func genEntity() domain.Room {
	return domain.Room{
		ID:    genID(),
		Title: genTitle(),
	}
}

func genID() domain.ID {
	return domain.ID(tdULID.Room.ID)
}

func genTitle() domain.Title {
	return domain.Title(tdString.Room.Title.Valid)
}
