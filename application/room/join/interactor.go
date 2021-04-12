package message

import (
	"github.com/oklog/ulid"

	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

type interactor struct {
	repository messageDomain.IRepository
}

// NewInteractor トークルーム入室アプリケーションサービスのコンストラクタ
func NewInteractor(r messageDomain.IRepository) IInputPort {
	return &interactor{
		repository: r,
	}
}

// Handle トークルーム入室アプリケーションサービス
func (i interactor) Handle(input InputData) {

	// RoomIDの生成
	parsedULID, err := ulid.Parse(input.RoomID)
	if err != nil {
		input.ErrCh <- err
	}
	roomID, err := roomDomain.NewID(&parsedULID)
	if err != nil {
		input.ErrCh <- err
	}

	// メッセージ一覧を返却
	messages, err := i.repository.FindAll(roomID)
	if err != nil {
		input.ErrCh <- err
	}
	for _, v := range *messages {
		input.MessageCh <- v
	}

	// 新規メッセージの監視
	if err = i.repository.Subscribe(input.Context, roomID, input.MessageCh); err != nil {
		input.ErrCh <- err
	}
}
