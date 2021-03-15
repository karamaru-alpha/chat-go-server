package room

import "errors"

// Room トークルームを表現するエンティティ
type Room struct {
	ID    ID
	Title Title
}

// NewRoom Roomエンティティを構築するコンストラクタ
func NewRoom(id *ID, title *Title) (*Room, error) {

	if id == nil {
		return nil, errors.New("RoomID is null")
	}
	if title == nil {
		return nil, errors.New("RoomTitle is null")
	}

	return &Room{ID: *id, Title: *title}, nil
}
