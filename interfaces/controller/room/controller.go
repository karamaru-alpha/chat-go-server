package room

import (
	"context"

	createApplication "github.com/karamaru-alpha/chat-go-server/application/room/create"
	findAllApplication "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	joinApplication "github.com/karamaru-alpha/chat-go-server/application/room/join"
	messageDTO "github.com/karamaru-alpha/chat-go-server/interfaces/dto/message"
	roomDTO "github.com/karamaru-alpha/chat-go-server/interfaces/dto/room"
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"
)

type controller struct {
	createApplication  createApplication.IInputPort
	findAllApplication findAllApplication.IInputPort
	joinApplication    joinApplication.IInputPort
}

// NewController gRPC経由のリクエストを捌くControllerのコンストラクタ
func NewController(
	c createApplication.IInputPort,
	f findAllApplication.IInputPort,
	j joinApplication.IInputPort,
) pb.RoomServicesServer {
	return &controller{
		createApplication:  c,
		findAllApplication: f,
		joinApplication:    j,
	}
}

// CreateRoom トークルームを作成するController
func (c controller) CreateRoom(ctx context.Context, request *pb.CreateRoomRequest) (*pb.CreateRoomResponse, error) {
	input := createApplication.InputData{Title: request.Title}

	output := c.createApplication.Handle(input)
	if output.Err != nil {
		return nil, output.Err
	}

	return &pb.CreateRoomResponse{Room: roomDTO.ToProto(output.Room)}, nil
}

// GetRooms トークルーム全件取得のController
func (c controller) GetRooms(ctx context.Context, _ *pb.GetRoomsRequest) (*pb.GetRoomsResponse, error) {
	output := c.findAllApplication.Handle()
	if output.Err != nil {
		return nil, output.Err
	}

	return &pb.GetRoomsResponse{Rooms: roomDTO.ToProtos(output.Rooms)}, nil
}

func (c controller) JoinRoom(ctx context.Context, request *pb.JoinRoomRequest) (*pb.JoinRoomResponse, error) {
	input := joinApplication.InputData{RoomID: request.RoomId}

	output := c.joinApplication.Handle(input)
	if output.Err != nil {
		return nil, output.Err
	}

	return &pb.JoinRoomResponse{Messages: messageDTO.ToProtos(output.Messages)}, nil
}
