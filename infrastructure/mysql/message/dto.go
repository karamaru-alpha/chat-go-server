package message

import (
	"github.com/oklog/ulid"

	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// Message メッセージエンティティにDB属性を加えたDTO
type Message struct {
	ID     string `gorm:"primary_key"`
	RoomID string
	Body   string
}

// ToDTO メッセージエンティティをDB情報を持った構造体に変換する
func ToDTO(entity messageDomain.Message) Message {
	if entity == (messageDomain.Message{}) {
		return Message{}
	}

	return Message{
		ID:     ulid.ULID(entity.ID).String(),
		RoomID: ulid.ULID(entity.RoomID).String(),
		Body:   string(entity.Body),
	}
}

// ToEntity DB情報を持った構造体からメッセージエンティティに変換する
func ToEntity(dto Message) (messageDomain.Message, error) {
	if dto == (Message{}) {
		return messageDomain.Message{}, nil
	}

	parsedMessageULID, err := ulid.Parse(dto.ID)
	if err != nil {
		return messageDomain.Message{}, err
	}
	parsedRoomULID, err := ulid.Parse(dto.RoomID)
	if err != nil {
		return messageDomain.Message{}, err
	}

	return messageDomain.Message{
		ID:     messageDomain.ID(parsedMessageULID),
		RoomID: roomDomain.ID(parsedRoomULID),
		Body:   messageDomain.Body(dto.Body),
	}, nil
}

// ToEntity DB情報を持った構造体のスライスをエンティティに変換する
func ToEntities(dtos []Message) ([]messageDomain.Message, error) {
	if len(dtos) == 0 {
		return nil, nil
	}

	entities := make([]messageDomain.Message, 0, len(dtos))
	for _, v := range dtos {
		entity, err := ToEntity(v)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}
