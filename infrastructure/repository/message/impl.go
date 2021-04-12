package message

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/oklog/ulid"

	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	mysql "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/message"
)

type repositoryImpl struct {
	gormDB      *gorm.DB
	redisClient *redis.Client
}

// NewRepositoryImpl メッセージの永続化・再構築を行うRepositoryImplのコンストラクタ
func NewRepositoryImpl(g *gorm.DB, r *redis.Client) messageDomain.IRepository {
	return &repositoryImpl{
		gormDB:      g,
		redisClient: r,
	}
}

// Save メッセージの登録を行う
func (r repositoryImpl) Save(ctx context.Context, entity *messageDomain.Message) error {
	dto := mysql.ToDTO(entity)
	if err := r.gormDB.Create(dto).Error; err != nil {
		return err
	}

	serializedMessage, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	r.redisClient.Publish(ctx, ulid.ULID(entity.RoomID).String(), serializedMessage)

	return nil
}

// FindAll 特定トークルームのメッセージ一覧を取得する
func (r repositoryImpl) FindAll(roomID *roomDomain.ID) (*[]messageDomain.Message, error) {
	var dtos []mysql.Message
	if err := r.gormDB.Where("room_id = ?", ulid.ULID(*roomID).String()).Find(&dtos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return mysql.ToEntities(&dtos)
}

// Subscribe 特定トークルームの新規メッセージを監視・再構築する
func (r repositoryImpl) Subscribe(
	ctx context.Context,
	roomID *roomDomain.ID,
	reciever chan messageDomain.Message,
) error {
	pubsub := r.redisClient.Subscribe(ctx, ulid.ULID(*roomID).String())
	defer pubsub.Close()
	if _, err := pubsub.Receive(ctx); err != nil {
		return err
	}

	ch := pubsub.Channel()
	for v := range ch {
		var latestMessage messageDomain.Message
		if err := json.Unmarshal([]byte(v.Payload), &latestMessage); err != nil {
			return err
		}
		reciever <- latestMessage
	}

	return nil
}
