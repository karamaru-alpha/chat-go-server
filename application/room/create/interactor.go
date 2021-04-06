package room

import (
	"errors"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	domainService "github.com/karamaru-alpha/chat-go-server/domain/service/room"
)

type interactor struct {
	factory       domainModel.IFactory
	repository    domainModel.IRepository
	domainService domainService.IDomainService
}

// NewInteractor トークルームを新規作成するアプリケーションサービスのコンストラクタ
func NewInteractor(f domainModel.IFactory, r domainModel.IRepository, s domainService.IDomainService) IInputPort {
	return &interactor{
		factory:       f,
		repository:    r,
		domainService: s,
	}
}

// Handle トークルームを新規作成するアプリケーションサービス
func (i interactor) Handle(input InputData) OutputData {

	roomTitle, err := domainModel.NewTitle(input.Title)
	if err != nil {
		return OutputData{Err: err}
	}

	room, err := i.factory.Create(roomTitle)
	if err != nil {
		return OutputData{Err: err}
	}

	isDup, err := i.domainService.Exists(room)
	if err != nil {
		return OutputData{Err: err}
	}
	if isDup {
		return OutputData{Err: errors.New("RoomTitle is duplicated")}
	}

	if err = i.repository.Save(room); err != nil {
		return OutputData{Err: err}
	}

	return OutputData{Room: room}
}
