package room

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	application "github.com/karamaru-alpha/chat-go-server/application/room/send_message"
	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"

	mockMessageDomain "github.com/karamaru-alpha/chat-go-server/mock/domain/model/message"
	mockRoomDomain "github.com/karamaru-alpha/chat-go-server/mock/domain/model/room"
	tdMessageDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdRoomDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdCommonString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/common"
	tdMessageString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/message"
)

type testHandler struct {
	interactor application.IInputPort

	factory           *mockMessageDomain.MockIFactory
	messageRepository *mockMessageDomain.MockIRepository
	roomRepository    *mockRoomDomain.MockIRepository
}

// TestHandle トークルームルームでメッセージを送信するアプリケーションサービスのテスト
func TestHandle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		before   func(testHandler)
		input    application.InputData
		expected application.OutputData
	}{
		{
			title: "【正常系】トークルームでメッセージを送信する",
			before: func(h testHandler) {
				h.roomRepository.EXPECT().Find(tdRoomDomain.ID).Return(tdRoomDomain.Entity, nil)
				h.factory.EXPECT().Create(tdRoomDomain.Entity, tdMessageDomain.Body).Return(tdMessageDomain.Entity, nil)
				h.messageRepository.EXPECT().Save(context.TODO(), tdMessageDomain.Entity).Return(nil)
			},
			input: application.InputData{
				Context: context.TODO(),
				RoomID:  tdCommonString.ULID.Valid,
				Body:    tdMessageString.Body.Valid,
			},
			expected: application.OutputData{
				Message: tdMessageDomain.Entity,
				Err:     nil,
			},
		},
		{
			title: "【異常系】本文が短い(空)",
			before: func(h testHandler) {
				h.roomRepository.EXPECT().Find(tdRoomDomain.ID).Return(tdRoomDomain.Entity, nil)
			},
			input: application.InputData{
				Context: context.TODO(),
				RoomID:  tdCommonString.ULID.Valid,
				Body:    tdMessageString.Body.Empty,
			},
			expected: application.OutputData{
				Message: messageDomain.Message{},
				Err:     errors.New("MessageBody is empty"),
			},
		},
		{
			title: "【異常系】本文が長い",
			before: func(h testHandler) {
				h.roomRepository.EXPECT().Find(tdRoomDomain.ID).Return(tdRoomDomain.Entity, nil)
			},
			input: application.InputData{
				Context: context.TODO(),
				RoomID:  tdCommonString.ULID.Valid,
				Body:    tdMessageString.Body.TooLong,
			},
			expected: application.OutputData{
				Message: messageDomain.Message{},
				Err:     errors.New("MessageBody should be 1 to 255 characters"),
			},
		},
		{
			title: "【異常系】ルームIDが空",
			input: application.InputData{
				Context: context.TODO(),
				RoomID:  "",
				Body:    tdMessageString.Body.Valid,
			},
			expected: application.OutputData{
				Message: messageDomain.Message{},
				Err:     errors.New("ulid: bad data size when unmarshaling"),
			},
		},
		{
			title: "【異常系】ルームIDが不正値",
			input: application.InputData{
				Context: context.TODO(),
				RoomID:  tdCommonString.ULID.Invalid,
				Body:    tdMessageString.Body.Valid,
			},
			expected: application.OutputData{
				Message: messageDomain.Message{},
				Err:     errors.New("ulid: bad data size when unmarshaling"),
			},
		},
		{
			title: "【異常系】存在しないルームID",
			before: func(h testHandler) {
				h.roomRepository.EXPECT().Find(tdRoomDomain.ID).Return(roomDomain.Room{}, nil)
				h.factory.EXPECT().Create(roomDomain.Room{}, tdMessageDomain.Body).Return(messageDomain.Message{}, errors.New("MessageRoom is not exist"))
			},
			input: application.InputData{
				Context: context.TODO(),
				RoomID:  tdCommonString.ULID.Valid,
				Body:    tdMessageString.Body.Valid,
			},
			expected: application.OutputData{
				Message: messageDomain.Message{},
				Err:     errors.New("MessageRoom is not exist"),
			},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Handle:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)

			if td.before != nil {
				td.before(tester)
			}

			output := tester.interactor.Handle(td.input)

			assert.Equal(t, td.expected, output)
		})
	}
}

func (h *testHandler) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	h.factory = mockMessageDomain.NewMockIFactory(ctrl)
	h.messageRepository = mockMessageDomain.NewMockIRepository(ctrl)
	h.roomRepository = mockRoomDomain.NewMockIRepository(ctrl)

	h.interactor = application.NewInteractor(
		h.factory,
		h.messageRepository,
		h.roomRepository,
	)
}
