package room

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	application "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	mockDomainModel "github.com/karamaru-alpha/chat-go-server/mock/domain/model/room"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
)

// TestHandle トークルームを全件取得するアプリケーションサービスのテスト
func TestHandle(t *testing.T) {
	t.Parallel()

	// go-mockの開始
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// reposityをモック
	repository := mockDomainModel.NewMockIRepository(ctrl)
	interactor := application.NewInteractor(repository)

	tests := []struct {
		title    string
		before   func()
		expected application.OutputData
	}{
		{
			title: "【正常系】",
			before: func() {
				repository.EXPECT().FindAll().Return(&[]domainModel.Room{tdDomain.Room.Entity}, nil)
			},
			expected: application.OutputData{
				Rooms: &[]domainModel.Room{tdDomain.Room.Entity},
				Err:   nil,
			},
		},
	}

	for _, td := range tests {
		td := td

		td.before()

		t.Run("Handle:"+td.title, func(t *testing.T) {
			output := interactor.Handle()
			assert.Equal(t, td.expected, output)
		})
	}

}
