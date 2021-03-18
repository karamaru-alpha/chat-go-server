package room

import (
	"github.com/oklog/ulid"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// Room トークルームエンティティにDB属性を加えたDTO
type Room struct {
	ID    string `gorm:"primary_key"`
	Title string
}

func ToDTO(entity *domainModel.Room) *Room {
	return &Room{
		ID:    ulid.ULID(entity.ID).String(),
		Title: string(entity.Title),
	}
}

func ToEntity(dto *Room) (*domainModel.Room, error) {
	parsedULID, err := ulid.Parse(dto.ID)
	if err != nil {
		return nil, err
	}

	entityID, err := domainModel.NewID(&parsedULID)
	if err != nil {
		return nil, err
	}

	entityTitle, err := domainModel.NewTitle(dto.Title)
	if err != nil {
		return nil, err
	}

	return domainModel.NewRoom(
		entityID,
		entityTitle,
	)
}

func ToEntities(dtos *[]Room) (*[]domainModel.Room, error) {
	var entities []domainModel.Room

	for _, v := range *dtos {
		entity, err := ToEntity(&v)
		if err != nil {
			return nil, err
		}
		entities = append(entities, *entity)
	}

	return &entities, nil
}
