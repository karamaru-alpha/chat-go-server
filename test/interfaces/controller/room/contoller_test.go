package room_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	createApplication "github.com/karamaru-alpha/chat-go-server/application/room/create"
	findAllApplication "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	joinApplication "github.com/karamaru-alpha/chat-go-server/application/room/join"
	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	controller "github.com/karamaru-alpha/chat-go-server/interfaces/controller/room"
	messageDTO "github.com/karamaru-alpha/chat-go-server/interfaces/dto/message"
	roomDTO "github.com/karamaru-alpha/chat-go-server/interfaces/dto/room"
	mockCreateApplication "github.com/karamaru-alpha/chat-go-server/mock/application/room/create"
	mockFindAllApplication "github.com/karamaru-alpha/chat-go-server/mock/application/room/find_all"
	mockJoinApplication "github.com/karamaru-alpha/chat-go-server/mock/application/room/join"
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"
	tdMessageDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdRoomDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

type controllerTester struct {
	controller         pb.RoomServicesServer
	createApplication  *mockCreateApplication.MockIInputPort
	findAllApplication *mockFindAllApplication.MockIInputPort
	joinApplication    *mockJoinApplication.MockIInputPort
}

// TestGetRooms トークルーム一覧取得Controllerのテスト
func TestGetRooms(t *testing.T) {
	t.Parallel()

	var tester controllerTester
	tester.setupTest(t)

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
				tester.findAllApplication.EXPECT().Handle().Return(findAllApplication.OutputData{
					Rooms: &[]roomDomain.Room{tdRoomDomain.Room.Entity}, Err: nil,
				})
			},
			input1:    context.TODO(),
			input2:    &pb.GetRoomsRequest{},
			expected1: &pb.GetRoomsResponse{Rooms: roomDTO.ToProtos(&[]roomDomain.Room{tdRoomDomain.Room.Entity})},
			expected2: nil,
		},
		{
			title: "【正常系】トークルームがまだない",
			before: func() {
				tester.findAllApplication.EXPECT().Handle().Return(findAllApplication.OutputData{
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
			t.Parallel()

			td.before()

			output1, output2 := tester.controller.GetRooms(td.input1, td.input2)

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

	var tester controllerTester
	tester.setupTest(t)

	tests := []struct {
		title     string
		before    func()
		input     *pb.CreateRoomRequest
		expected1 *pb.CreateRoomResponse
		expected2 error
	}{
		{
			title: "【正常系】トークルーム作成",
			before: func() {
				tester.createApplication.EXPECT().Handle(
					createApplication.InputData{Title: tdString.Room.Title.Valid},
				).Return(
					createApplication.OutputData{Room: &tdRoomDomain.Room.Entity, Err: nil},
				)
			},
			input:     &pb.CreateRoomRequest{Title: tdString.Room.Title.Valid},
			expected1: &pb.CreateRoomResponse{Room: roomDTO.ToProto(&tdRoomDomain.Room.Entity)},
			expected2: nil,
		},
		{
			title: "【異常系】タイトルが不正値(empty)",
			before: func() {
				tester.createApplication.EXPECT().Handle(
					createApplication.InputData{Title: ""},
				).Return(
					createApplication.OutputData{Room: nil, Err: errors.New("error")},
				)
			},
			input:     &pb.CreateRoomRequest{Title: ""},
			expected1: nil,
			expected2: errors.New("error"),
		},
		{
			title: "【異常系】タイトルが不正値(long)",
			before: func() {
				tester.createApplication.EXPECT().Handle(
					createApplication.InputData{Title: tdString.Room.Title.TooLong},
				).Return(
					createApplication.OutputData{Room: nil, Err: errors.New("error")},
				)
			},
			input:     &pb.CreateRoomRequest{Title: tdString.Room.Title.TooLong},
			expected1: nil,
			expected2: errors.New("error"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("CreateRoom:"+td.title, func(t *testing.T) {
			t.Parallel()

			td.before()

			output1, output2 := tester.controller.CreateRoom(context.TODO(), td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}

// TestJoinRoom トークルーム入室Controllerのテスト
func TestJoinRoom(t *testing.T) {
	t.Parallel()

	var tester controllerTester
	tester.setupTest(t)

	tests := []struct {
		title     string
		before    func()
		input     *pb.JoinRoomRequest
		expected1 *pb.JoinRoomResponse
		expected2 error
	}{
		{
			title: "【正常系】トークルーム入室",
			before: func() {
				tester.joinApplication.EXPECT().Handle(
					joinApplication.InputData{RoomID: tdString.Room.ID.Valid},
				).Return(
					joinApplication.OutputData{
						Messages: &[]messageDomain.Message{tdMessageDomain.Message.Entity}, Err: nil,
					},
				)
			},
			input: &pb.JoinRoomRequest{RoomId: tdString.Room.ID.Valid},
			expected1: &pb.JoinRoomResponse{
				Messages: messageDTO.ToProtos(&[]messageDomain.Message{tdMessageDomain.Message.Entity}),
			},
			expected2: nil,
		},
		{
			title: "【異常系】RoomIDが不正値",
			before: func() {
				tester.joinApplication.EXPECT().Handle(
					joinApplication.InputData{RoomID: tdString.Room.ID.Invalid},
				).Return(
					joinApplication.OutputData{Messages: nil, Err: errors.New("error")},
				)
			},
			input:     &pb.JoinRoomRequest{RoomId: tdString.Room.ID.Invalid},
			expected1: nil,
			expected2: errors.New("error"),
		},
		{
			title: "【異常系】RoomIDが空",
			before: func() {
				tester.joinApplication.EXPECT().Handle(
					joinApplication.InputData{RoomID: ""},
				).Return(
					joinApplication.OutputData{Messages: nil, Err: errors.New("error")},
				)
			},
			input:     &pb.JoinRoomRequest{RoomId: ""},
			expected1: nil,
			expected2: errors.New("error"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("JoinRoom:"+td.title, func(t *testing.T) {
			t.Parallel()

			td.before()

			output1, output2 := tester.controller.JoinRoom(context.TODO(), td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}

func (c *controllerTester) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	c.createApplication = mockCreateApplication.NewMockIInputPort(ctrl)
	c.findAllApplication = mockFindAllApplication.NewMockIInputPort(ctrl)
	c.joinApplication = mockJoinApplication.NewMockIInputPort(ctrl)
	c.controller = controller.NewController(c.createApplication, c.findAllApplication, c.joinApplication)
}
