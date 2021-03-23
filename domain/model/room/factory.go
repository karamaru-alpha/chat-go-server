package room

import (
	"github.com/oklog/ulid"
)

// iFactory Roomエンティティの生成処理を担うFactoryのインターフェース
type iFactory interface {
	Create(*Title) (*Room, error)
}

type factory struct {
	generateULID func() ulid.ULID
}

// NewFactory Roomエンティティの生成処理を担うFactoryのコンストラクタ
func NewFactory(generateULID func() ulid.ULID) iFactory {
	return &factory{
		generateULID,
	}
}

// Create Roomエンティティの生成処理を担うファクトリ
func (f factory) Create(title *Title) (*Room, error) {

	ulid := f.generateULID()

	roomID, err := NewID(&ulid)
	if err != nil {
		return nil, err
	}

	return NewRoom(roomID, title)
}
