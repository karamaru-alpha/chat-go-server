package room

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	application "github.com/karamaru-alpha/chat-go-server/application/room/find_all"
	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	mockDomainModel "github.com/karamaru-alpha/chat-go-server/mock/domain/model/room"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain"
)

// TestHandle トークルームを全件取得するアプリケーションサービスのテスト
func TestHandle(t *testing.T) {
	t.Parallel()

	// go-mockの開始
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// reposityをモック
	repository := mockDomainModel.NewMockIRepository(ctrl)
	repository.EXPECT().FindAll().Return(&[]domainModel.Room{tdDomain.Room.Entity.Valid}, nil)

	interactor := application.NewInteractor(repository)

	tests := []struct {
		title    string
		expected application.OutputData
	}{
		{
			title: "【正常系】",
			expected: application.OutputData{
				Rooms: &[]domainModel.Room{tdDomain.Room.Entity.Valid},
				Err:   nil,
			},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Handle:"+td.title, func(t *testing.T) {
			output := interactor.Handle()
			assert.Equal(t, td.expected, output)
		})
	}

}