package room

import (
	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// IDomainService ドメインサービスのインターフェース
type IDomainService interface {
	Exists(domainModel.Room) (bool, error)
}

type domainService struct {
	repository domainModel.IRepository
}

// NewDomainService ドメインサービスのコンストラクタ
func NewDomainService(r domainModel.IRepository) IDomainService {
	return &domainService{
		repository: r,
	}
}

// Exists トークルーム重複判定のドメインサービス
func (s domainService) Exists(room domainModel.Room) (bool, error) {

	room, err := s.repository.FindByTitle(room.Title)

	return room != domainModel.Room{}, err
}
