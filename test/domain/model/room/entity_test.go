package room

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// TestNewRoom トークルームコンストラクタのテスト
func TestNewRoom(t *testing.T) {
	t.Parallel()

	roomID := domainModel.ID(testData.id.valid)
	roomTitle := domainModel.Title(testData.title.valid)

	tests := []struct {
		title     string
		input1    *domainModel.ID
		input2    *domainModel.Title
		expected1 *domainModel.Room
		expected2 error
	}{
		{
			title:     "【正常系】",
			input1:    &roomID,
			input2:    &roomTitle,
			expected1: &domainModel.Room{ID: roomID, Title: roomTitle},
			expected2: nil,
		},
		{
			title:     "【異常系】IDがnil",
			input1:    nil,
			input2:    &roomTitle,
			expected1: &domainModel.Room{ID: roomID, Title: roomTitle},
			expected2: errors.New("RoomID is null"),
		},
		{
			title:     "【異常系】Titleがnil",
			input1:    &roomID,
			input2:    nil,
			expected1: &domainModel.Room{ID: roomID, Title: roomTitle},
			expected2: errors.New("RoomTitle is null"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("NewRoom:"+td.title, func(t *testing.T) {

			output1, output2 := domainModel.NewRoom(td.input1, td.input2)

			assert.IsType(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
