package room

import (
	"errors"

	"github.com/jinzhu/gorm"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

type repositoryImpl struct {
	db *gorm.DB
}

// NewRepositoryImpl リポジトリを実現するimplを生成するコンストラクタ
func NewRepositoryImpl(db *gorm.DB) domainModel.IRepository {
	return &repositoryImpl{
		db,
	}
}

// Save トークルームをDBに保存し永続化させる
func (r repositoryImpl) Save(entity *domainModel.Room) error {
	dto := ToDTO(entity)
	return r.db.Create(dto).Error
}

// FindAll トークルーム一覧をDBから再構築する
func (r repositoryImpl) FindAll() (*[]domainModel.Room, error) {
	var dtos []Room

	if err := r.db.Find(&dtos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return ToEntities(&dtos)
}
