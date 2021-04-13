package room

import (
	"errors"

	"github.com/oklog/ulid"
)

// ID トークルームの識別子を表現する値オブジェクト
type ID ulid.ULID

// NewID トークルームIDの値オブジェクトを生成するコンストラクタ
func NewID(v ulid.ULID) (ID, error) {
	if v == (ulid.ULID{}) {
		return ID{}, errors.New("RoomID is null")
	}

	return ID(v), nil
}
