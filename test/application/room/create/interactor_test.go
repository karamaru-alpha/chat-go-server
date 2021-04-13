package room

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	application "github.com/karamaru-alpha/chat-go-server/application/room/create"
	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"

	mockDomain "github.com/karamaru-alpha/chat-go-server/mock/domain/model/room"
	mockDomainService "github.com/karamaru-alpha/chat-go-server/mock/domain/service/room"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/room"
)

type testHandler struct {
	interactor application.IInputPort

	factory       *mockDomain.MockIFactory
	repository    *mockDomain.MockIRepository
	domainService *mockDomainService.MockIDomainService
}

// TestHandle トークルームを作成するアプリケーションサービスのテスト
func TestHandle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		before   func(testHandler)
		input    application.InputData
		expected application.OutputData
	}{
		{
			title: "【正常系】トークルーム作成",
			before: func(h testHandler) {
				h.factory.EXPECT().Create(tdDomain.Title).Return(tdDomain.Entity, nil)
				h.domainService.EXPECT().Exists(tdDomain.Entity).Return(false, nil)
				h.repository.EXPECT().Save(tdDomain.Entity).Return(nil)
			},
			input: application.InputData{
				Title: tdString.Title.Valid,
			},
			expected: application.OutputData{
				Room: tdDomain.Entity,
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
				Title: tdString.Title.TooShort,
			},
			expected: application.OutputData{
				Room: domain.Room{},
				Err:  errors.New("RoomTitle should be Three to twenty characters"),
			},
		},
		{
			title: "【異常系】タイトルが長い",
			input: application.InputData{
				Title: tdString.Title.TooLong,
			},
			expected: application.OutputData{
				Room: domain.Room{},
				Err:  errors.New("RoomTitle should be Three to twenty characters"),
			},
		},
		{
			title: "【異常系】タイトルが重複している",
			before: func(h testHandler) {
				h.factory.EXPECT().Create(tdDomain.Title).Return(tdDomain.Entity, nil)
				h.domainService.EXPECT().Exists(tdDomain.Entity).Return(true, nil)
			},
			input: application.InputData{
				Title: tdString.Title.Valid,
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
	h.factory = mockDomain.NewMockIFactory(ctrl)
	h.repository = mockDomain.NewMockIRepository(ctrl)
	h.domainService = mockDomainService.NewMockIDomainService(ctrl)

	h.interactor = application.NewInteractor(
		h.factory,
		h.repository,
		h.domainService,
	)
}
