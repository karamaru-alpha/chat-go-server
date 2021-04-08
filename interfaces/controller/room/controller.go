package room

import (
	"context"

	createApplication "github.com/karamaru-alpha/chat-go-server/application/room/create"
	findAllApplication "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"
)

type controller struct {
	createApplication  createApplication.IInputPort
	findAllApplication findAllApplication.IInputPort
}

// NewController gRPC経由のリクエストを捌くControllerのコンストラクタ
func NewController(c createApplication.IInputPort, f findAllApplication.IInputPort) pb.RoomServicesServer {
	return &controller{
		createApplication:  c,
		findAllApplication: f,
	}
}

// CreateRoom トークルームを作成するController
func (h controller) CreateRoom(ctx context.Context, request *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {

	input := createApplication.InputData{Title: request.Title}
	output := h.createApplication.Handle(input)

	if output.Err != nil {
		return nil, output.Err
	}

	return &pb.CreateRoomResponse{Room: ToProto(output.Room)}, nil
}

// GetRooms トークルーム全件取得のController
func (h controller) GetRooms(ctx context.Context, _ *pb.GetRoomsRequest) (*pb.GetRoomsResponse, error) {

	output := h.findAllApplication.Handle()

	if output.Err != nil {
		return nil, output.Err
	}

	return &pb.GetRoomsResponse{Rooms: ToProtos(output.Rooms)}, nil
}
