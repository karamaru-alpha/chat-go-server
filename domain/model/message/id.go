package message

import (
	"errors"

	"github.com/oklog/ulid"
)

// ID メッセージ識別子を表す値オブジェクト
type ID ulid.ULID

// NewID メッセージ識別子を表す値オブジェクトのコンストラクタ
func NewID(v *ulid.ULID) (*ID, error) {

	if v == nil {
		return nil, errors.New("MessageID is null")
	}

	id := ID(*v)
	return &id, nil
}
