package testdata

import (
	"log"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

// Message メッセージエンティティのテストデータ
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
	factory := domainModel.NewFactory(mockUtil.NewULIDGenerator())

	roomID := genRoomID()
	body := genBody()
	message, err := factory.Create(&roomID, &body)
	if err != nil {
		log.Fatal(err)
	}

	return *message
}

func genID() domainModel.ID {
	id, err := domainModel.NewID(&tdULID.Message.ID)
	if err != nil {
		log.Fatal(err)
	}

	return *id
}

func genRoomID() domainModel.RoomID {
	roomID, err := domainModel.NewRoomID(&tdULID.Room.ID)
	if err != nil {
		log.Fatal(err)
	}

	return *roomID
}

func genBody() domainModel.Body {
	body, err := domainModel.NewBody(tdString.Message.Body.Valid)
	if err != nil {
		log.Fatal(err)
	}

	return *body
}
