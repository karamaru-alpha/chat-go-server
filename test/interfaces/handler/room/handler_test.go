package room

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	createApplication "github.com/karamaru-alpha/chat-go-server/application/room/create"
	findAllApplication "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	dto "github.com/karamaru-alpha/chat-go-server/interfaces/dto/room"
	handler "github.com/karamaru-alpha/chat-go-server/interfaces/handler/room"
	pb "github.com/karamaru-alpha/chat-go-server/interfaces/proto/pb"
	mockCreateApplication "github.com/karamaru-alpha/chat-go-server/mock/application/room/create"
	mockFindAllApplication "github.com/karamaru-alpha/chat-go-server/mock/application/room/find_all"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

// TestGetRooms トークルーム一覧取得Handlerのテスト
func TestGetRooms(t *testing.T) {
	t.Parallel()

	// go-mockの開始
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// applicationをモック
	createUsecase := mockCreateApplication.NewMockIInputPort(ctrl)
	findAllUsecase := mockFindAllApplication.NewMockIInputPort(ctrl)
	findAllUsecase.EXPECT().Handle().Return(
		findAllApplication.OutputData{
			Rooms: &[]domainModel.Room{tdDomain.Room.Entity.Valid}, Err: nil,
		},
	)

	controller := handler.NewHandler(createUsecase, findAllUsecase)

	tests := []struct {
		title     string
		expected1 *pb.GetRoomsResponse
		expected2 error
	}{
		{
			title:     "【正常系】",
			expected1: &pb.GetRoomsResponse{Rooms: dto.ToProtos(&[]domainModel.Room{tdDomain.Room.Entity.Valid})},
			expected2: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("GetRooms:"+td.title, func(t *testing.T) {
			output1, output2 := controller.GetRooms(context.TODO(), &pb.GetRoomsRequest{})
			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}

// TestCreateRoom トークルーム作成Handlerのテスト
func TestCreateRoom(t *testing.T) {
	t.Parallel()

	// go-mockの開始
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// applicationをモック
	createUsecase := mockCreateApplication.NewMockIInputPort(ctrl)
	findAllUsecase := mockFindAllApplication.NewMockIInputPort(ctrl)
	createUsecase.EXPECT().Handle(
		createApplication.InputData{Title: tdString.Room.Title.Valid},
	).Return(
		createApplication.OutputData{Room: &tdDomain.Room.Entity.Valid, Err: nil},
	)

	controller := handler.NewHandler(createUsecase, findAllUsecase)

	tests := []struct {
		title     string
		input     *pb.CreateRoomRequest
		expected1 *pb.CreateRoomResponse
		expected2 error
	}{
		{
			title:     "【正常系】",
			input:     &pb.CreateRoomRequest{Title: tdString.Room.Title.Valid},
			expected1: &pb.CreateRoomResponse{Room: dto.ToProto(&tdDomain.Room.Entity.Valid)},
			expected2: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("CreateRoom:"+td.title, func(t *testing.T) {
			output1, output2 := controller.CreateRoom(context.TODO(), td.input)
			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
