package message

import (
	"errors"

	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	"github.com/karamaru-alpha/chat-go-server/util"
)

// IFactory メッセージファクトリのインターフェース
type IFactory interface {
	Create(roomDomain.Room, Body) (Message, error)
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
func (f factory) Create(room roomDomain.Room, body Body) (Message, error) {
	if room == (roomDomain.Room{}) {
		return Message{}, errors.New("MessageRoom is not exist")
	}

	messageID, err := NewID(f.ulidGenerator.Generate())
	if err != nil {
		return Message{}, err
	}

	return NewMessage(messageID, room.ID, body)
}
