package room_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	createApplication "github.com/karamaru-alpha/chat-go-server/application/room/create"
	findAllApplication "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	controller "github.com/karamaru-alpha/chat-go-server/interfaces/controller/room"
	mockCreateApplication "github.com/karamaru-alpha/chat-go-server/mock/application/room/create"
	mockFindAllApplication "github.com/karamaru-alpha/chat-go-server/mock/application/room/find_all"
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

// TestGetRooms トークルーム一覧取得Controllerのテスト
func TestGetRooms(t *testing.T) {
	t.Parallel()

	// go-mockの開始
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// applicationをモック
	createUsecase := mockCreateApplication.NewMockIInputPort(ctrl)
	findAllUsecase := mockFindAllApplication.NewMockIInputPort(ctrl)

	handler := controller.NewController(createUsecase, findAllUsecase)

	tests := []struct {
		title     string
		before    func()
		input1    context.Context
		input2    *pb.GetRoomsRequest
		expected1 *pb.GetRoomsResponse
		expected2 error
	}{
		{
			title: "【正常系】トークルームが1つ",
			before: func() {
				findAllUsecase.EXPECT().Handle().Return(findAllApplication.OutputData{
					Rooms: &[]domain.Room{tdDomain.Room.Entity}, Err: nil,
				})
			},
			input1:    context.TODO(),
			input2:    &pb.GetRoomsRequest{},
			expected1: &pb.GetRoomsResponse{Rooms: controller.ToProtos(&[]domain.Room{tdDomain.Room.Entity})},
			expected2: nil,
		},
		{
			title: "【正常系】トークルームがまだない",
			before: func() {
				findAllUsecase.EXPECT().Handle().Return(findAllApplication.OutputData{
					Rooms: nil, Err: nil,
				})
			},
			input1:    context.TODO(),
			input2:    &pb.GetRoomsRequest{},
			expected1: &pb.GetRoomsResponse{Rooms: nil},
			expected2: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("GetRooms:"+td.title, func(t *testing.T) {
			td.before()

			output1, output2 := handler.GetRooms(td.input1, td.input2)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}

// TestCreateRoom トークルーム作成Controllerのテスト
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
		createApplication.OutputData{Room: &tdDomain.Room.Entity, Err: nil},
	)

	handler := controller.NewController(createUsecase, findAllUsecase)

	tests := []struct {
		title     string
		input     *pb.CreateRoomRequest
		expected1 *pb.CreateRoomResponse
		expected2 error
	}{
		{
			title:     "【正常系】",
			input:     &pb.CreateRoomRequest{Title: tdString.Room.Title.Valid},
			expected1: &pb.CreateRoomResponse{Room: controller.ToProto(&tdDomain.Room.Entity)},
			expected2: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("CreateRoom:"+td.title, func(t *testing.T) {
			output1, output2 := handler.CreateRoom(context.TODO(), td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
