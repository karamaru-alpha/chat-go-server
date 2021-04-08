package message

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	application "github.com/karamaru-alpha/chat-go-server/application/room/join"
	domain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	mockDomain "github.com/karamaru-alpha/chat-go-server/mock/domain/model/message"
	tdMessageDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdRoomDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

type interactorTester struct {
	interactor application.IInputPort
	repository *mockDomain.MockIRepository
}

// TestHandle トークルーム入室アプリケーションサービスのテスト
func TestHandle(t *testing.T) {
	t.Parallel()

	var tester interactorTester
	tester.setupTest(t)

	tests := []struct {
		title    string
		before   func()
		input    application.InputData
		expected application.OutputData
	}{
		{
			title: "【正常系】トークルームのメッセージが1件",
			before: func() {
				tester.repository.EXPECT().FindAll(&tdRoomDomain.Room.ID).Return(
					&[]domain.Message{tdMessageDomain.Message.Entity}, nil,
				)
			},
			input: application.InputData{RoomID: tdString.Room.ID.Valid},
			expected: application.OutputData{
				Messages: &[]domain.Message{tdMessageDomain.Message.Entity},
				Err:      nil,
			},
		},
		{
			title: "【正常系】該当トークルームのメッセージが存在しない",
			before: func() {
				tester.repository.EXPECT().FindAll(&tdRoomDomain.Room.ID).Return(nil, nil)
			},
			input: application.InputData{RoomID: tdString.Room.ID.Valid},
			expected: application.OutputData{
				Messages: nil,
				Err:      nil,
			},
		},
		{
			title: "【異常系】ルームIDが空",
			input: application.InputData{
				RoomID: "",
			},
			expected: application.OutputData{
				Messages: nil,
				Err:      errors.New("ulid: bad data size when unmarshaling"),
			},
		},
		{
			title: "【異常系】トークルームのIDが不正値",
			input: application.InputData{RoomID: tdString.Room.ID.Invalid},
			expected: application.OutputData{
				Messages: nil,
				Err:      errors.New("ulid: bad data size when unmarshaling"),
			},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Handle:"+td.title, func(t *testing.T) {
			t.Parallel()

			if td.before != nil {
				td.before()
			}

			output := tester.interactor.Handle(td.input)

			assert.Equal(t, td.expected, output)
		})
	}
}

func (i *interactorTester) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	i.repository = mockDomain.NewMockIRepository(ctrl)
	i.interactor = application.NewInteractor(i.repository)
}
