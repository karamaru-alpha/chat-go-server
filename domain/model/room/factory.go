package room

import (
	"github.com/oklog/ulid"
)

// IFactory Roomエンティティの生成処理を担うFactoryのインターフェース
type IFactory interface {
	Create(*Title) (*Room, error)
}

type factory struct {
	generateULID func() ulid.ULID
}

// NewFactory Roomエンティティの生成処理を担うFactoryのコンストラクタ
func NewFactory(generateULID func() ulid.ULID) IFactory {
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
