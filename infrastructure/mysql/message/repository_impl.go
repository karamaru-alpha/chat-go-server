package message

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/oklog/ulid"

	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

type repositoryImpl struct {
	db *gorm.DB
}

// NewRepositoryImpl メッセージの永続化・再構築を行うRepositoryImplのコンストラクタ
func NewRepositoryImpl(db *gorm.DB) messageDomain.IRepository {
	return &repositoryImpl{
		db,
	}
}

// Create メッセージの永続化を行う
func (r repositoryImpl) Create(entity *messageDomain.Message) error {
	dto := ToDTO(entity)
	return r.db.Create(dto).Error
}

// FindAll 特定トークルームのメッセージ一覧を取得する
func (r repositoryImpl) FindAll(roomID *roomDomain.ID) (*[]messageDomain.Message, error) {
	var dtos []Message
	if err := r.db.Where("room_id = ?", ulid.ULID(*roomID).String()).Find(&dtos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return ToEntities(&dtos)
}
