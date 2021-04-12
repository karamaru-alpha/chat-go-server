package room

import (
	"context"

	createApplication "github.com/karamaru-alpha/chat-go-server/application/room/create"
	findAllApplication "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	joinApplication "github.com/karamaru-alpha/chat-go-server/application/room/join"
	sendMessageApplication "github.com/karamaru-alpha/chat-go-server/application/room/send_message"
	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	messageDTO "github.com/karamaru-alpha/chat-go-server/interfaces/dto/message"
	roomDTO "github.com/karamaru-alpha/chat-go-server/interfaces/dto/room"
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"
)

type controller struct {
	createApplication      createApplication.IInputPort
	findAllApplication     findAllApplication.IInputPort
	joinApplication        joinApplication.IInputPort
	sendMessageApplication sendMessageApplication.IInputPort
}

// NewController gRPC経由のリクエストを捌くControllerのコンストラクタ
func NewController(
	c createApplication.IInputPort,
	f findAllApplication.IInputPort,
	j joinApplication.IInputPort,
	s sendMessageApplication.IInputPort,
) pb.RoomServicesServer {
	return &controller{
		createApplication:      c,
		findAllApplication:     f,
		joinApplication:        j,
		sendMessageApplication: s,
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

// JoinRoom トークルーム入室のController
func (c controller) JoinRoom(request *pb.JoinRoomRequest, stream pb.RoomServices_JoinRoomServer) error {
	messageCh := make(chan messageDomain.Message)
	errCh := make(chan error)

	input := joinApplication.InputData{
		Context:   stream.Context(),
		RoomID:    request.RoomId,
		MessageCh: messageCh,
		ErrCh:     errCh,
	}

	go c.joinApplication.Handle(input)
	go func() {
		for v := range messageCh {
			err := stream.Send(&pb.JoinRoomResponse{
				Message: messageDTO.ToProto(&v),
			})

			if err != nil {
				errCh <- err
			}
		}
	}()

	return <-errCh
}

// SendMessage メッセージ送信のController
func (c controller) SendMessage(ctx context.Context, request *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	input := sendMessageApplication.InputData{Context: ctx, RoomID: request.RoomId, Body: request.Body}

	output := c.sendMessageApplication.Handle(input)
	if output.Err != nil {
		return nil, output.Err
	}

	return &pb.SendMessageResponse{}, nil
}
