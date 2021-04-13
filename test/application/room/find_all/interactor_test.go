package room

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	application "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"

	mockDomain "github.com/karamaru-alpha/chat-go-server/mock/domain/model/room"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
)

type testHandler struct {
	interactor application.IInputPort

	repository *mockDomain.MockIRepository
}

// TestHandle トークルームを全件取得するアプリケーションサービスのテスト
func TestHandle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		before   func(testHandler)
		expected application.OutputData
	}{
		{
			title: "【正常系】トークルーム全検索",
			before: func(h testHandler) {
				h.repository.EXPECT().FindAll().Return([]domain.Room{tdDomain.Entity}, nil)
			},
			expected: application.OutputData{
				Rooms: []domain.Room{tdDomain.Entity},
				Err:   nil,
			},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Handle:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)

			td.before(tester)

			output := tester.interactor.Handle()

			assert.Equal(t, td.expected, output)
		})
	}
}

func (h *testHandler) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	h.repository = mockDomain.NewMockIRepository(ctrl)

	h.interactor = application.NewInteractor(h.repository)
}
