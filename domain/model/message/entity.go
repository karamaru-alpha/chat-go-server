package message

import "errors"

// Message メッセージエンティティ
type Message struct {
	ID     ID
	RoomID RoomID
	Body   Body
}

// Message メッセージエンティティのコンストラクタ
func NewMessage(id *ID, roomID *RoomID, body *Body) (*Message, error) {

	if id == nil {
		return nil, errors.New("MessageID is null")
	}

	if roomID == nil {
		return nil, errors.New("MessageRoomID is null")
	}

	if body == nil {
		return nil, errors.New("MessageBody is null")
	}

	return &Message{RoomID: *roomID, Body: *body}, nil
}
