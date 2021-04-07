package message

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	application "github.com/karamaru-alpha/chat-go-server/application/message/create"
	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	mockMessageDomain "github.com/karamaru-alpha/chat-go-server/mock/domain/model/message"
	mockRoomDomain "github.com/karamaru-alpha/chat-go-server/mock/domain/model/room"
	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	tdMessageDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdRoomDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

// TestHandle トークルームを作成するアプリケーションサービスのテスト
func TestHandle(t *testing.T) {
	t.Parallel()

	// go-mockの開始
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// reposityをモック
	messageRepository := mockMessageDomain.NewMockIRepository(ctrl)
	roomRepository := mockRoomDomain.NewMockIRepository(ctrl)

	factory := messageDomain.NewFactory(mockUtil.NewULIDGenerator())
	interactor := application.NewInteractor(factory, messageRepository, roomRepository)

	tests := []struct {
		title    string
		before   func()
		input    application.InputData
		expected application.OutputData
	}{
		{
			title: "【正常系】",
			before: func() {
				roomRepository.EXPECT().Find(&tdRoomDomain.Room.ID).Return(&tdRoomDomain.Room.Entity, nil)
				messageRepository.EXPECT().Create(&tdMessageDomain.Message.Entity).Return(nil)
			},
			input: application.InputData{
				RoomID: tdString.Room.ID.Valid,
				Body:   tdString.Message.Body.Valid,
			},
			expected: application.OutputData{
				Message: &tdMessageDomain.Message.Entity,
				Err:     nil,
			},
		},
		{
			title: "【異常系】本文が短い(空)",
			before: func() {
				roomRepository.EXPECT().Find(&tdRoomDomain.Room.ID).Return(&tdRoomDomain.Room.Entity, nil)
			},
			input: application.InputData{
				RoomID: tdString.Room.ID.Valid,
				Body:   tdString.Message.Body.Empty,
			},
			expected: application.OutputData{
				Message: nil,
				Err:     errors.New("MessageBody is empty"),
			},
		},
		{
			title: "【異常系】本文が長い",
			before: func() {
				roomRepository.EXPECT().Find(&tdRoomDomain.Room.ID).Return(&tdRoomDomain.Room.Entity, nil)
			},
			input: application.InputData{
				RoomID: tdString.Room.ID.Valid,
				Body:   tdString.Message.Body.TooLong,
			},
			expected: application.OutputData{
				Message: nil,
				Err:     errors.New("MessageBody should be 1 to 255 characters"),
			},
		},
		{
			title: "【異常系】ルームIDが空",
			input: application.InputData{
				RoomID: "",
				Body:   tdString.Message.Body.Valid,
			},
			expected: application.OutputData{
				Message: nil,
				Err:     errors.New("ulid: bad data size when unmarshaling"),
			},
		},
		{
			title: "【異常系】ルームIDが不正値",
			input: application.InputData{
				RoomID: tdString.Room.ID.Invalid,
				Body:   tdString.Message.Body.Valid,
			},
			expected: application.OutputData{
				Message: nil,
				Err:     errors.New("ulid: bad data size when unmarshaling"),
			},
		},
		{
			title: "【異常系】存在しないルームID",
			before: func() {
				roomRepository.EXPECT().Find(&tdRoomDomain.Room.ID).Return(nil, nil)
			},
			input: application.InputData{
				RoomID: tdString.Room.ID.Valid,
				Body:   tdString.Message.Body.Valid,
			},
			expected: application.OutputData{
				Message: nil,
				Err:     errors.New("MessageRoom is not exist"),
			},
		},
	}

	for _, td := range tests {
		td := td

		if td.before != nil {
			td.before()
		}

		t.Run("Handle:"+td.title, func(t *testing.T) {
			output := interactor.Handle(td.input)
			assert.Equal(t, td.expected, output)
		})
	}

}
