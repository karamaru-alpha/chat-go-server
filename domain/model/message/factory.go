package message

import "github.com/karamaru-alpha/chat-go-server/util"

// IFactory メッセージファクトリのインターフェース
type IFactory interface {
	Create(*RoomID, *Body) (*Message, error)
}

type factory struct {
	ulidGenerator util.IULIDGenerator
}

// IFactory メッセージファクトリコンストラクタ
func NewFactory(ulidGenerator util.IULIDGenerator) IFactory {
	return &factory{
		ulidGenerator,
	}
}

// Create メッセージエンティティの生成処理を担うファクトリ
func (f factory) Create(roomID *RoomID, body *Body) (*Message, error) {

	ulid := f.ulidGenerator.Generate()

	messageID, err := NewID(&ulid)
	if err != nil {
		return nil, err
	}

	return NewMessage(messageID, roomID, body)
}
