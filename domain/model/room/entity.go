package room

import (
	"errors"
)

// Room トークルームを表現するエンティティ
type Room struct {
	ID    ID
	Title Title
}

// NewRoom Roomエンティティを構築するコンストラクタ
func NewRoom(id ID, title Title) (Room, error) {
	if id == (ID{}) {
		return Room{}, errors.New("RoomID is null")
	}
	if title == "" {
		return Room{}, errors.New("RoomTitle is null")
	}

	return Room{ID: id, Title: title}, nil
}
