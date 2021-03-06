package room

import (
	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

type interactor struct {
	repository domain.IRepository
}

// NewInteractor トークルームを全件取得するアプリケーションサービスのコンストラクタ
func NewInteractor(repository domain.IRepository) IInputPort {
	return &interactor{
		repository,
	}
}

// Handle トークルームを全件取得するアプリケーションサービス
func (i interactor) Handle() OutputData {
	rooms, err := i.repository.FindAll()
	return OutputData{
		rooms,
		err,
	}
}
