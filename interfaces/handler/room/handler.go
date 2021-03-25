package room

import (
	"context"

	createApplication "github.com/karamaru-alpha/chat-go-server/application/room/create"
	findAllApplication "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	dto "github.com/karamaru-alpha/chat-go-server/interfaces/dto/room"
	pb "github.com/karamaru-alpha/chat-go-server/interfaces/proto/pb"
)

type handler struct {
	createApplication  createApplication.IInputPort
	findAllApplication findAllApplication.IInputPort
}

// NewHandler gRPC経由のリクエストを捌くHandlerのコンストラクタ
func NewHandler(createApplication createApplication.IInputPort, findAllApplication findAllApplication.IInputPort) pb.RoomServicesServer {
	return &handler{
		createApplication,
		findAllApplication,
	}
}

// CreateRoom トークルームを作成するHandler
func (h handler) CreateRoom(ctx context.Context, request *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {

	input := createApplication.InputData{Title: request.Title}
	output := h.createApplication.Handle(input)

	if output.Err != nil {
		return nil, output.Err
	}

	return &pb.CreateRoomResponse{Room: dto.ToProto(output.Room)}, nil
}

// GetRooms トークルーム全件取得のHandler
func (h handler) GetRooms(ctx context.Context, _ *pb.GetRoomsRequest) (*pb.GetRoomsResponse, error) {

	output := h.findAllApplication.Handle()

	if output.Err != nil {
		return nil, output.Err
	}

	return &pb.GetRoomsResponse{Rooms: dto.ToProtos(output.Rooms)}, nil
}
