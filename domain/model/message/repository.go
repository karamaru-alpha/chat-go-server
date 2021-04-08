package message

import "github.com/karamaru-alpha/chat-go-server/domain/model/room"

// IRepository メッセージエンティティの永続化・再構築を実現するリポジトリ
type IRepository interface {
	Save(*Message) error
	FindAll(*room.ID) (*[]Message, error)
}
