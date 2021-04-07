package room

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/oklog/ulid"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

type repositoryImpl struct {
	db *gorm.DB
}

// NewRepositoryImpl リポジトリを実現するimplを生成するコンストラクタ
func NewRepositoryImpl(db *gorm.DB) domain.IRepository {
	return &repositoryImpl{
		db,
	}
}

// Save トークルームをDBに保存し永続化させる
func (r repositoryImpl) Save(entity *domain.Room) error {
	dto := ToDTO(entity)
	return r.db.Create(dto).Error
}

// FindAll トークルーム一覧をDBから再構築する
func (r repositoryImpl) FindAll() (*[]domain.Room, error) {
	var dtos []Room

	if err := r.db.Find(&dtos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return ToEntities(&dtos)
}

// Find トークルームをIDから１件検索・再構成する
func (r repositoryImpl) Find(id *domain.ID) (*domain.Room, error) {
	var dto Room

	if err := r.db.Where("id = ?", ulid.ULID(*id).String()).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return ToEntity(&dto)
}

// FindByTitle トークルームをタイトルから一件取得する
func (r repositoryImpl) FindByTitle(title *domain.Title) (*domain.Room, error) {
	var dto Room

	if err := r.db.Where("title = ?", string(*title)).First(&dto).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return ToEntity(&dto)
}
