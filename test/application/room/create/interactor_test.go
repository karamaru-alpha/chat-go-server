package room

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	application "github.com/karamaru-alpha/chat-go-server/application/room/create"
	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	domainService "github.com/karamaru-alpha/chat-go-server/domain/service/room"
	mockDomain "github.com/karamaru-alpha/chat-go-server/mock/domain/model/room"
	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

type interactorTester struct {
	interactor application.IInputPort
	repository *mockDomain.MockIRepository
}

// TestHandle トークルームを作成するアプリケーションサービスのテスト
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
			title: "【正常系】トークルーム作成",
			before: func() {
				tester.repository.EXPECT().Save(tdDomain.Room.Entity).Return(nil)
				tester.repository.EXPECT().FindByTitle(tdDomain.Room.Title).Return(domain.Room{}, nil)
			},
			input: application.InputData{
				Title: tdString.Room.Title.Valid,
			},
			expected: application.OutputData{
				Room: tdDomain.Room.Entity,
				Err:  nil,
			},
		},
		{
			title: "【異常系】タイトルが空文字列",
			input: application.InputData{
				Title: "",
			},
			expected: application.OutputData{
				Room: domain.Room{},
				Err:  errors.New("RoomTitle is null"),
			},
		},
		{
			title: "【異常系】タイトルが短い",
			input: application.InputData{
				Title: tdString.Room.Title.TooShort,
			},
			expected: application.OutputData{
				Room: domain.Room{},
				Err:  errors.New("RoomTitle should be Three to twenty characters"),
			},
		},
		{
			title: "【異常系】タイトルが長い",
			input: application.InputData{
				Title: tdString.Room.Title.TooLong,
			},
			expected: application.OutputData{
				Room: domain.Room{},
				Err:  errors.New("RoomTitle should be Three to twenty characters"),
			},
		},
		{
			title: "【異常系】タイトルが重複している",
			before: func() {
				tester.repository.EXPECT().FindByTitle(tdDomain.Room.Title).Return(tdDomain.Room.Entity, nil)
			},
			input: application.InputData{
				Title: tdString.Room.Title.Valid,
			},
			expected: application.OutputData{
				Room: domain.Room{},
				Err:  errors.New("RoomTitle is duplicated"),
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
	i.interactor = application.NewInteractor(
		domain.NewFactory(mockUtil.NewULIDGenerator()),
		i.repository,
		domainService.NewDomainService(i.repository),
	)
}
