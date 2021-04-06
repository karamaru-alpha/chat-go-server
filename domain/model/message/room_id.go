package message

import (
	"errors"

	"github.com/oklog/ulid"
)

// RoomID メッセージが紐づくトークルームの識別子を表す値オブジェクト
type RoomID ulid.ULID

// NewRoomID メッセージが紐づくトークルームの識別子を表す値オブジェクトコンストラクタ
func NewRoomID(v *ulid.ULID) (*RoomID, error) {

	if v == nil {
		return nil, errors.New("MessageRoomID is null")
	}

	roomID := RoomID(*v)
	return &roomID, nil
}
