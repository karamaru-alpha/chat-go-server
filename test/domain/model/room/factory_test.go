package room

import (
	"testing"

	"github.com/stretchr/testify/assert"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	testdata "github.com/karamaru-alpha/chat-go-server/test/testdata"
)

// TestCreate トークルーム生成処理を担うファクトリのテスト
func TestCreate(t *testing.T) {
	t.Parallel()

	roomTitle, err := domainModel.NewTitle(testdata.Room.Title.Valid)
	assert.NoError(t, err)

	factory := domainModel.NewFactory(mockUtil.GenerateULID)

	tests := []struct {
		title     string
		input     *domainModel.Title
		expected1 *domainModel.Room
		expected2 error
	}{
		{
			title:     "【正常系】",
			input:     roomTitle,
			expected1: &domainModel.Room{ID: domainModel.ID(testdata.Room.ID.Valid), Title: *roomTitle},
			expected2: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Create:"+td.title, func(t *testing.T) {

			output1, output2 := factory.Create(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
