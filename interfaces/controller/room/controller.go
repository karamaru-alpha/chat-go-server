package room

import (
	"context"

	"github.com/oklog/ulid"

	createApplication "github.com/karamaru-alpha/chat-go-server/application/room/create"
	findAllApplication "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	joinApplication "github.com/karamaru-alpha/chat-go-server/application/room/join"
	sendMessageApplication "github.com/karamaru-alpha/chat-go-server/application/room/send_message"
	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
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

	return &pb.CreateRoomResponse{
		Room: &pb.Room{
			Id:    ulid.ULID(output.Room.ID).String(),
			Title: string(output.Room.Title),
		},
	}, nil
}

// GetRooms トークルーム全件取得のController
func (c controller) GetRooms(ctx context.Context, _ *pb.GetRoomsRequest) (*pb.GetRoomsResponse, error) {
	output := c.findAllApplication.Handle()
	if output.Err != nil {
		return nil, output.Err
	}

	responseRooms := make([]*pb.Room, 0, len(output.Rooms))
	for _, v := range output.Rooms {
		responseRooms = append(responseRooms, &pb.Room{
			Id:    ulid.ULID(v.ID).String(),
			Title: string(v.Title),
		})
	}
	return &pb.GetRoomsResponse{Rooms: responseRooms}, nil
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
				Message: &pb.Message{
					Id:     ulid.ULID(v.ID).String(),
					RoomId: ulid.ULID(v.ID).String(),
					Body:   string(v.Body),
				},
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
