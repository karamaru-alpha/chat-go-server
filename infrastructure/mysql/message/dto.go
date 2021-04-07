package message

import (
	"github.com/oklog/ulid"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// Message メッセージエンティティにDB属性を加えたDTO
type Message struct {
	ID     string `gorm:"primary_key"`
	RoomID string
	Body   string
}

// ToDTO メッセージエンティティをDB情報を持った構造体に変換する
func ToDTO(entity *domain.Message) *Message {
	if entity == nil {
		return nil
	}

	return &Message{
		ID:     ulid.ULID(entity.ID).String(),
		RoomID: ulid.ULID(entity.RoomID).String(),
		Body:   string(entity.Body),
	}
}

// ToEntity DB情報を持った構造体からメッセージエンティティに変換する
func ToEntity(dto *Message) (*domain.Message, error) {
	if dto == nil {
		return nil, nil
	}

	parsedULID, err := ulid.Parse(dto.ID)
	if err != nil {
		return nil, err
	}

	entityID, err := domain.NewID(&parsedULID)
	if err != nil {
		return nil, err
	}

	parsedRoomULID, err := ulid.Parse(dto.RoomID)
	if err != nil {
		return nil, err
	}

	entityRoomID, err := roomDomain.NewID(&parsedRoomULID)
	if err != nil {
		return nil, err
	}

	entityBody, err := domain.NewBody(dto.Body)
	if err != nil {
		return nil, err
	}

	return domain.NewMessage(
		entityID,
		entityRoomID,
		entityBody,
	)
}

// ToEntity DB情報を持った構造体のスライスをエンティティに変換する
func ToEntities(dtos *[]Message) (*[]domain.Message, error) {
	if dtos == nil {
		return nil, nil
	}

	entities := make([]domain.Message, 0, len(*dtos))
	for _, v := range *dtos {
		entity, err := ToEntity(&v)
		if err != nil {
			return nil, err
		}
		entities = append(entities, *entity)
	}
	return &entities, nil
}
