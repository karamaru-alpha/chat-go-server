package message

import (
	"github.com/oklog/ulid"

	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

type interactor struct {
	repository messageDomain.IRepository
}

// NewInteractor メッセージ一覧取得アプリケーションサービスのコンストラクタ
func NewInteractor(r messageDomain.IRepository) IInputPort {
	return &interactor{
		repository: r,
	}
}

// Handle メッセージ一覧取得アプリケーションサービス
func (i interactor) Handle(input InputData) OutputData {

	parsedULID, err := ulid.Parse(input.RoomID)
	if err != nil {
		return OutputData{Err: err}
	}

	roomID, err := roomDomain.NewID(&parsedULID)
	if err != nil {
		return OutputData{Err: err}
	}

	messages, err := i.repository.FindAll(roomID)
	if err != nil {
		return OutputData{Err: err}
	}

	return OutputData{Messages: messages}
}
