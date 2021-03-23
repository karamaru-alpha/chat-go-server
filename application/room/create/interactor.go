package room

import (
	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

type interactor struct {
	factory    domainModel.IFactory
	repository domainModel.IRepository
}

// NewInteractor トークルームを新規作成するアプリケーションサービスのコンストラクタ
func NewInteractor(factory domainModel.IFactory, repository domainModel.IRepository) IInputPort {
	return &interactor{
		factory,
		repository,
	}
}

// Handle トークルームを新規作成するアプリケーションサービス
func (i interactor) Handle(input InputData) OutputData {

	roomTitle, err := domainModel.NewTitle(input.Title)
	if err != nil {
		return OutputData{nil, err}
	}

	room, err := i.factory.Create(roomTitle)
	if err != nil {
		return OutputData{nil, err}
	}

	if err = i.repository.Save(room); err != nil {
		return OutputData{nil, err}
	}

	return OutputData{room, nil}
}
