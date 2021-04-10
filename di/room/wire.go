// +build wireinject

package room

import (
	"github.com/google/wire"

	createApplication "github.com/karamaru-alpha/chat-go-server/application/room/create"
	findAllApplication "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	joinApplication "github.com/karamaru-alpha/chat-go-server/application/room/join"
	sendMessageApplication "github.com/karamaru-alpha/chat-go-server/application/room/send_message"
	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	domainService "github.com/karamaru-alpha/chat-go-server/domain/service/room"
	"github.com/karamaru-alpha/chat-go-server/infrastructure/mysql"
	messageRepositoryImpl "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/message"
	roomRepositoryImpl "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/room"
	controller "github.com/karamaru-alpha/chat-go-server/interfaces/controller/room"
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"
	"github.com/karamaru-alpha/chat-go-server/util"
)

// DI dependency injection about room
func DI() pb.RoomServicesServer {
	wire.Build(
		controller.NewController,
		createApplication.NewInteractor,
		findAllApplication.NewInteractor,
		joinApplication.NewInteractor,
		sendMessageApplication.NewInteractor,
		roomRepositoryImpl.NewRepositoryImpl,
		messageRepositoryImpl.NewRepositoryImpl,
		roomDomain.NewFactory,
		messageDomain.NewFactory,
		domainService.NewDomainService,
		util.NewULIDGenerator,
		mysql.ConnectGorm,
	)

	return nil
}
