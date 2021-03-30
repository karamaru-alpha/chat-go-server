// +build wireinject

package room

import (
	"github.com/google/wire"

	createApplication "github.com/karamaru-alpha/chat-go-server/application/room/create"
	findAllApplication "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	"github.com/karamaru-alpha/chat-go-server/infrastructure/mysql"
	repositoryImpl "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/room"
	handler "github.com/karamaru-alpha/chat-go-server/interfaces/handler/room"
	pb "github.com/karamaru-alpha/chat-go-server/interfaces/proto/pb"
	"github.com/karamaru-alpha/chat-go-server/util"
)

// DI dependency injection about room
func DI() pb.RoomServicesServer {
	wire.Build(
		handler.NewHandler,
		createApplication.NewInteractor,
		findAllApplication.NewInteractor,
		repositoryImpl.NewRepositoryImpl,
		domainModel.NewFactory,
		util.NewULIDGenerator,
		mysql.ConnectGorm,
	)

	return nil
}
