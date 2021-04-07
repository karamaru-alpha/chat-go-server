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

func ToDTO(entity *domain.Room) *Room {
	return &Room{
		ID:    ulid.ULID(entity.ID).String(),
		Title: string(entity.Title),
	}
}

func ToEntity(dto *Room) (*domain.Room, error) {
	parsedULID, err := ulid.Parse(dto.ID)
	if err != nil {
		return nil, err
	}

	entityID, err := domain.NewID(&parsedULID)
	if err != nil {
		return nil, err
	}

	entityTitle, err := domain.NewTitle(dto.Title)
	if err != nil {
		return nil, err
	}

	return domain.NewRoom(
		entityID,
		entityTitle,
	)
}

func ToEntities(dtos *[]Room) (*[]domain.Room, error) {
	entities := make([]domain.Room, 0, len(*dtos))
	for _, v := range *dtos {
		entity, err := ToEntity(&v)
		if err != nil {
			return nil, err
		}
		entities = append(entities, *entity)
	}
	return &entities, nil
}
