package room_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	createApplication "github.com/karamaru-alpha/chat-go-server/application/room/create"
	findAllApplication "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	sendMessageApplication "github.com/karamaru-alpha/chat-go-server/application/room/send_message"
	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	controller "github.com/karamaru-alpha/chat-go-server/interfaces/controller/room"
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"

	mockCreateApplication "github.com/karamaru-alpha/chat-go-server/mock/application/room/create"
	mockFindAllApplication "github.com/karamaru-alpha/chat-go-server/mock/application/room/find_all"
	mockJoinApplication "github.com/karamaru-alpha/chat-go-server/mock/application/room/join"
	mockSendMessageApplication "github.com/karamaru-alpha/chat-go-server/mock/application/room/send_message"
	tdMessageDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdRoomDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdRoomPb "github.com/karamaru-alpha/chat-go-server/test/testdata/pb/room"
	tdCommonString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/common"
	tdMessageString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/message"
	tdRoomString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/room"
)

type testHandler struct {
	controller pb.RoomServicesServer

	context                context.Context
	createApplication      *mockCreateApplication.MockIInputPort
	findAllApplication     *mockFindAllApplication.MockIInputPort
	joinApplication        *mockJoinApplication.MockIInputPort
	sendMessageApplication *mockSendMessageApplication.MockIInputPort
}

// TestGetRooms トークルーム一覧取得Controllerのテスト
func TestGetRooms(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		before    func(testHandler)
		input     *pb.GetRoomsRequest
		expected1 *pb.GetRoomsResponse
		expected2 error
	}{
		{
			title: "【正常系】トークルームが1つ",
			before: func(h testHandler) {
				h.findAllApplication.EXPECT().Handle().Return(findAllApplication.OutputData{
					Rooms: []roomDomain.Room{tdRoomDomain.Entity}, Err: nil,
				})
			},
			input:     &pb.GetRoomsRequest{},
			expected1: &pb.GetRoomsResponse{Rooms: []*pb.Room{&tdRoomPb.Room}},
			expected2: nil,
		},
		{
			title: "【正常系】トークルームがまだない",
			before: func(h testHandler) {
				h.findAllApplication.EXPECT().Handle().Return(findAllApplication.OutputData{
					Rooms: nil, Err: nil,
				})
			},
			input:     &pb.GetRoomsRequest{},
			expected1: &pb.GetRoomsResponse{Rooms: []*pb.Room{}},
			expected2: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("GetRooms:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)

			td.before(tester)

			output1, output2 := tester.controller.GetRooms(tester.context, td.input)

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

	tests := []struct {
		title     string
		before    func(testHandler)
		input     *pb.CreateRoomRequest
		expected1 *pb.CreateRoomResponse
		expected2 error
	}{
		{
			title: "【正常系】トークルーム作成",
			before: func(h testHandler) {
				h.createApplication.EXPECT().Handle(
					createApplication.InputData{Title: tdRoomString.Title.Valid},
				).Return(
					createApplication.OutputData{Room: tdRoomDomain.Entity, Err: nil},
				)
			},
			input:     &pb.CreateRoomRequest{Title: tdRoomString.Title.Valid},
			expected1: &pb.CreateRoomResponse{Room: &tdRoomPb.Room},
			expected2: nil,
		},
		{
			title: "【異常系】タイトルが不正値(empty)",
			before: func(h testHandler) {
				h.createApplication.EXPECT().Handle(
					createApplication.InputData{Title: ""},
				).Return(
					createApplication.OutputData{Room: roomDomain.Room{}, Err: errors.New("error")},
				)
			},
			input:     &pb.CreateRoomRequest{Title: ""},
			expected1: nil,
			expected2: errors.New("error"),
		},
		{
			title: "【異常系】タイトルが不正値(long)",
			before: func(h testHandler) {
				h.createApplication.EXPECT().Handle(
					createApplication.InputData{Title: tdRoomString.Title.TooLong},
				).Return(
					createApplication.OutputData{Room: roomDomain.Room{}, Err: errors.New("error")},
				)
			},
			input:     &pb.CreateRoomRequest{Title: tdRoomString.Title.TooLong},
			expected1: nil,
			expected2: errors.New("error"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("CreateRoom:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)

			td.before(tester)

			output1, output2 := tester.controller.CreateRoom(tester.context, td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}

// TODO
// TestJoinRoom トークルーム入室Controllerのテスト
func TestJoinRoom(t *testing.T) {
	t.Parallel()

	assert.Equal(t, true, true)
}

// TestSendMessage トークルームでメッセージ送信するControllerのテスト
func TestSendMessage(t *testing.T) {
	t.Parallel()

	var tester testHandler
	tester.setupTest(t)

	tests := []struct {
		title     string
		before    func(testHandler)
		input     *pb.SendMessageRequest
		expected1 *pb.SendMessageResponse
		expected2 error
	}{
		{
			title: "【正常系】メッセージ送信",
			before: func(h testHandler) {
				h.sendMessageApplication.EXPECT().Handle(
					sendMessageApplication.InputData{
						Context: h.context,
						RoomID:  tdCommonString.ULID.Valid,
						Body:    tdMessageString.Body.Valid,
					},
				).Return(
					sendMessageApplication.OutputData{Message: tdMessageDomain.Entity, Err: nil},
				)
			},
			input:     &pb.SendMessageRequest{RoomId: tdCommonString.ULID.Valid, Body: tdMessageString.Body.Valid},
			expected1: &pb.SendMessageResponse{},
			expected2: nil,
		},
		{
			title: "【異常系】メッセージが長すぎる",
			before: func(h testHandler) {
				h.sendMessageApplication.EXPECT().Handle(
					sendMessageApplication.InputData{
						Context: h.context,
						RoomID:  tdCommonString.ULID.Valid,
						Body:    tdMessageString.Body.TooLong,
					},
				).Return(
					sendMessageApplication.OutputData{Message: messageDomain.Message{}, Err: errors.New("error")},
				)
			},
			input:     &pb.SendMessageRequest{RoomId: tdCommonString.ULID.Valid, Body: tdMessageString.Body.TooLong},
			expected1: nil,
			expected2: errors.New("error"),
		},
		{
			title: "【異常系】メッセージが空",
			before: func(h testHandler) {
				h.sendMessageApplication.EXPECT().Handle(
					sendMessageApplication.InputData{
						Context: h.context,
						RoomID:  tdCommonString.ULID.Valid,
						Body:    tdMessageString.Body.Empty,
					},
				).Return(
					sendMessageApplication.OutputData{Message: messageDomain.Message{}, Err: errors.New("error")},
				)
			},
			input:     &pb.SendMessageRequest{RoomId: tdCommonString.ULID.Valid, Body: tdMessageString.Body.Empty},
			expected1: nil,
			expected2: errors.New("error"),
		},
		{
			title: "【異常系】RoomIDが不正値",
			before: func(h testHandler) {
				h.sendMessageApplication.EXPECT().Handle(
					sendMessageApplication.InputData{
						Context: h.context,
						RoomID:  tdCommonString.ULID.Invalid,
						Body:    tdMessageString.Body.Valid,
					},
				).Return(
					sendMessageApplication.OutputData{Message: messageDomain.Message{}, Err: errors.New("error")},
				)
			},
			input:     &pb.SendMessageRequest{RoomId: tdCommonString.ULID.Invalid, Body: tdMessageString.Body.Valid},
			expected1: nil,
			expected2: errors.New("error"),
		},
		{
			title: "【異常系】RoomIDが空",
			before: func(h testHandler) {
				h.sendMessageApplication.EXPECT().Handle(
					sendMessageApplication.InputData{
						Context: h.context,
						RoomID:  "",
						Body:    tdMessageString.Body.Valid,
					},
				).Return(
					sendMessageApplication.OutputData{Message: messageDomain.Message{}, Err: errors.New("error")},
				)
			},
			input:     &pb.SendMessageRequest{RoomId: "", Body: tdMessageString.Body.Valid},
			expected1: nil,
			expected2: errors.New("error"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("SendMessage:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			td.before(tester)

			output1, output2 := tester.controller.SendMessage(tester.context, td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}

func (c *testHandler) setupTest(t *testing.T) {
	c.context = context.TODO()

	ctrl := gomock.NewController(t)
	c.createApplication = mockCreateApplication.NewMockIInputPort(ctrl)
	c.findAllApplication = mockFindAllApplication.NewMockIInputPort(ctrl)
	c.joinApplication = mockJoinApplication.NewMockIInputPort(ctrl)
	c.sendMessageApplication = mockSendMessageApplication.NewMockIInputPort(ctrl)

	c.controller = controller.NewController(
		c.createApplication,
		c.findAllApplication,
		c.joinApplication,
		c.sendMessageApplication,
	)
}
