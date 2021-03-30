package testdata

import (
	"log"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

var Room = struct {
	Entity entity
	ID     id
	Title  title
}{
	Entity: entity{
		Valid: genRoom(),
	},
	ID: id{
		Valid: genRoomID(),
	},
	Title: title{
		Valid: genRoomTitle(),
	},
}

type entity struct {
	Valid domainModel.Room
}

type id struct {
	Valid domainModel.ID
}

type title struct {
	Valid domainModel.Title
}

func genRoom() domainModel.Room {
	factory := domainModel.NewFactory(mockUtil.NewULIDGenerator())

	roomTitle := genRoomTitle()
	room, err := factory.Create(&roomTitle)
	if err != nil {
		log.Fatal(err)
	}

	return *room
}

func genRoomID() domainModel.ID {
	id, err := domainModel.NewID(&tdULID.Room.ID.Valid)
	if err != nil {
		log.Fatal(err)
	}

	return *id
}

func genRoomTitle() domainModel.Title {
	title, err := domainModel.NewTitle(tdString.Room.Title.Valid)
	if err != nil {
		log.Fatal(err)
	}

	return *title
}
