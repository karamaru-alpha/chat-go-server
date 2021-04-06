package testdata

import (
	"log"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

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
	factory := domainModel.NewFactory(mockUtil.NewULIDGenerator())

	roomTitle := genTitle()
	room, err := factory.Create(&roomTitle)
	if err != nil {
		log.Fatal(err)
	}

	return *room
}

func genID() domainModel.ID {
	id, err := domainModel.NewID(&tdULID.Room.ID)
	if err != nil {
		log.Fatal(err)
	}

	return *id
}

func genTitle() domainModel.Title {
	title, err := domainModel.NewTitle(tdString.Room.Title.Valid)
	if err != nil {
		log.Fatal(err)
	}

	return *title
}
