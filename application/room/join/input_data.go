package message

import (
	"context"

	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
)

// InputData トークルーム入室の入力データ
type InputData struct {
	Context   context.Context
	RoomID    string
	MessageCh chan messageDomain.Message
	ErrCh     chan error
}
