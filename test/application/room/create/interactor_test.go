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
	"github.com/karamaru-alpha/chat-go-server/test/testdata"
)

// TestHandle トークルームを作成するアプリケーションサービスのテスト
func TestHandle(t *testing.T) {
	t.Parallel()

	// go-mockの開始
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 正常なリクエストが来た場合に永続化したいRoomEntityを作成
	roomTitle, err := domainModel.NewTitle(testdata.Room.Title.Valid)
	assert.NoError(t, err)
	factory := domainModel.NewFactory(mockUtil.GenerateULID)
	room, err := factory.Create(roomTitle)
	assert.NoError(t, err)

	// reposityをモック
	repository := mockDomainModel.NewMockIRepository(ctrl)
	repository.EXPECT().Save(room).Return(nil)

	interactor := application.NewInteractor(factory, repository)

	tests := []struct {
		title    string
		input    application.InputData
		expected application.OutputData
	}{
		{
			title: "【正常系】",
			input: application.InputData{
				Title: testdata.Room.Title.Valid,
			},
			expected: application.OutputData{
				Room: room,
				Err:  nil,
			},
		},
		{
			title: "【異常系】タイトルが短い",
			input: application.InputData{
				Title: testdata.Room.Title.TooShort,
			},
			expected: application.OutputData{
				Room: nil,
				Err:  errors.New("RoomTitle should be Three to twenty characters"),
			},
		},
		{
			title: "【異常系】タイトルが長い",
			input: application.InputData{
				Title: testdata.Room.Title.TooLong,
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
