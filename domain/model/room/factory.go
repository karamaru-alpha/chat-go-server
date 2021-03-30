package room

import (
	"github.com/karamaru-alpha/chat-go-server/mock/util"
)

// IFactory Roomエンティティの生成処理を担うFactoryのインターフェース
type IFactory interface {
	Create(*Title) (*Room, error)
}

type factory struct {
	ulidGenerator util.IULIDGenerator
}

// NewFactory Roomエンティティの生成処理を担うFactoryのコンストラクタ
func NewFactory(ulidGenerator util.IULIDGenerator) IFactory {
	return &factory{
		ulidGenerator,
	}
}

// Create Roomエンティティの生成処理を担うファクトリ
func (f factory) Create(title *Title) (*Room, error) {

	ulid := f.ulidGenerator.Generate()

	roomID, err := NewID(&ulid)
	if err != nil {
		return nil, err
	}

	return NewRoom(roomID, title)
}
