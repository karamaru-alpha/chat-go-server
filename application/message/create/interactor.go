package message

import (
	"github.com/oklog/ulid"

	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

type interactor struct {
	factory           messageDomain.IFactory
	messageRepository messageDomain.IRepository
	roomRepository    roomDomain.IRepository
}

// NewInteractor メッセージ作成アプリケーションサービスのコンストラクタ
func NewInteractor(
	f messageDomain.IFactory,
	mr messageDomain.IRepository,
	rr roomDomain.IRepository,
) IInputPort {
	return &interactor{
		factory:           f,
		messageRepository: mr,
		roomRepository:    rr,
	}
}

// Handle メッセージ作成アプリケーションサービス
func (i interactor) Handle(input InputData) OutputData {

	parsedULID, err := ulid.Parse(input.RoomID)
	if err != nil {
		return OutputData{Err: err}
	}
	roomID, err := roomDomain.NewID(&parsedULID)
	if err != nil {
		return OutputData{Err: err}
	}

	room, err := i.roomRepository.Find(roomID)
	if err != nil {
		return OutputData{Err: err}
	}

	messageBody, err := messageDomain.NewBody(input.Body)
	if err != nil {
		return OutputData{Err: err}
	}

	message, err := i.factory.Create(room, messageBody)
	if err != nil {
		return OutputData{Err: err}
	}

	if err = i.messageRepository.Save(message); err != nil {
		return OutputData{Err: err}
	}

	return OutputData{Message: message}
}
