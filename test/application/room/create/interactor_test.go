package room

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	application "github.com/karamaru-alpha/chat-go-server/application/room/create"
	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	mockDomainModel "github.com/karamaru-alpha/chat-go-server/mock/domain/model/room"
	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

// TestHandle トークルームを作成するアプリケーションサービスのテスト
func TestHandle(t *testing.T) {
	t.Parallel()

	// go-mockの開始
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// reposityをモック
	repository := mockDomainModel.NewMockIRepository(ctrl)
	repository.EXPECT().Save(&tdDomain.Room.Entity).Return(nil)

	factory := domainModel.NewFactory(mockUtil.NewULIDGenerator())
	interactor := application.NewInteractor(factory, repository)

	tests := []struct {
		title    string
		input    application.InputData
		expected application.OutputData
	}{
		{
			title: "【正常系】",
			input: application.InputData{
				Title: tdString.Room.Title.Valid,
			},
			expected: application.OutputData{
				Room: &tdDomain.Room.Entity,
				Err:  nil,
			},
		},
		{
			title: "【異常系】タイトルが短い",
			input: application.InputData{
				Title: tdString.Room.Title.TooShort,
			},
			expected: application.OutputData{
				Room: nil,
				Err:  errors.New("RoomTitle should be Three to twenty characters"),
			},
		},
		{
			title: "【異常系】タイトルが長い",
			input: application.InputData{
				Title: tdString.Room.Title.TooLong,
			},
			expected: application.OutputData{
				Room: nil,
				Err:  errors.New("RoomTitle should be Three to twenty characters"),
			},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Handle:"+td.title, func(t *testing.T) {
			output := interactor.Handle(td.input)
			assert.Equal(t, td.expected, output)
		})
	}

}
