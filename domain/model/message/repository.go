package message

import (
	"context"

	"github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// IRepository メッセージエンティティの永続化・再構築を実現するリポジトリ
type IRepository interface {
	Save(context.Context, Message) error
	FindAll(room.ID) ([]Message, error)
	Subscribe(context.Context, room.ID, chan Message) error
}
