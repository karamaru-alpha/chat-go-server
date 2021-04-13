package room

import (
	"github.com/oklog/ulid"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// Room トークルームエンティティにDB属性を加えたDTO
type Room struct {
	ID    string `gorm:"primary_key"`
	Title string
}

// ToDTO トークルームエンティティをDB情報を持ったDTOに変換する
func ToDTO(entity domain.Room) Room {
	if entity == (domain.Room{}) {
		return Room{}
	}

	return Room{
		ID:    ulid.ULID(entity.ID).String(),
		Title: string(entity.Title),
	}
}

// ToEntity DB情報を持った構造体をトークルームエンティティに変換する
func ToEntity(dto Room) (domain.Room, error) {
	if dto == (Room{}) {
		return domain.Room{}, nil
	}

	parsedULID, err := ulid.Parse(dto.ID)
	if err != nil {
		return domain.Room{}, err
	}

	return domain.Room{
		ID:    domain.ID(parsedULID),
		Title: domain.Title(dto.Title),
	}, nil
}

// ToEntities DB情報を持った複数の構造体をトークルームエンティティに変換する
func ToEntities(dtos []Room) ([]domain.Room, error) {
	if len(dtos) == 0 {
		return nil, nil
	}

	entities := make([]domain.Room, 0, len(dtos))
	for _, v := range dtos {
		entity, err := ToEntity(v)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}
