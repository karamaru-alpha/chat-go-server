// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package room

import (
	room4 "github.com/karamaru-alpha/chat-go-server/application/room/create"
	room5 "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	message2 "github.com/karamaru-alpha/chat-go-server/application/room/join"
	"github.com/karamaru-alpha/chat-go-server/domain/model/room"
	room3 "github.com/karamaru-alpha/chat-go-server/domain/service/room"
	"github.com/karamaru-alpha/chat-go-server/infrastructure/mysql"
	"github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/message"
	room2 "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/room"
	room6 "github.com/karamaru-alpha/chat-go-server/interfaces/controller/room"
	"github.com/karamaru-alpha/chat-go-server/proto/pb"
	"github.com/karamaru-alpha/chat-go-server/util"
)

// Injectors from wire.go:

// DI dependency injection about room
func DI() proto.RoomServicesServer {
	iulidGenerator := util.NewULIDGenerator()
	iFactory := room.NewFactory(iulidGenerator)
	db := mysql.ConnectGorm()
	iRepository := room2.NewRepositoryImpl(db)
	iDomainService := room3.NewDomainService(iRepository)
	iInputPort := room4.NewInteractor(iFactory, iRepository, iDomainService)
	roomIInputPort := room5.NewInteractor(iRepository)
	messageIRepository := message.NewRepositoryImpl(db)
	messageIInputPort := message2.NewInteractor(messageIRepository)
	roomServicesServer := room6.NewController(iInputPort, roomIInputPort, messageIInputPort)
	return roomServicesServer
}
