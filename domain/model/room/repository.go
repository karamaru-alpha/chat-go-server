package room

// IRepository トークルームを永続化・再構築するリポジトリ
type IRepository interface {
	Save(*Room) error
	FindAll() (*[]Room, error)
	FindByTitle(*Title) (*Room, error)
}
