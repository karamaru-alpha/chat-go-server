package message

// IRepository メッセージエンティティの永続化・再構築を実現するリポジトリ
type IRepository interface {
	Create(*Message) error
	FindAll(*RoomID) (*[]Message, error)
}
