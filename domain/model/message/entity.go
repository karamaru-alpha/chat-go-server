package message

import (
	"errors"

	"github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// Message メッセージエンティティ
type Message struct {
	ID     ID
	RoomID room.ID
	Body   Body
}

// Message メッセージエンティティのコンストラクタ
func NewMessage(id *ID, roomID *room.ID, body *Body) (*Message, error) {

	if id == nil {
		return nil, errors.New("MessageID is null")
	}

	if roomID == nil {
		return nil, errors.New("MessageRoomID is null")
	}

	if body == nil {
		return nil, errors.New("MessageBody is null")
	}

	return &Message{ID: *id, RoomID: *roomID, Body: *body}, nil
}
