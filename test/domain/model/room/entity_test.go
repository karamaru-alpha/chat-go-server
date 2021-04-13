package room

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
)

// TestNewRoom トークルームコンストラクタのテスト
func TestNewRoom(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		input1    domain.ID
		input2    domain.Title
		expected1 domain.Room
		expected2 error
	}{
		{
			title:     "【正常系】トークルームエンティティを生成",
			input1:    tdDomain.Room.ID,
			input2:    tdDomain.Room.Title,
			expected1: tdDomain.Room.Entity,
			expected2: nil,
		},
		{
			title:     "【異常系】IDがゼロ値",
			input1:    domain.ID{},
			input2:    tdDomain.Room.Title,
			expected1: domain.Room{},
			expected2: errors.New("RoomID is null"),
		},
		{
			title:     "【異常系】Titleが空文字列",
			input1:    tdDomain.Room.ID,
			input2:    "",
			expected1: domain.Room{},
			expected2: errors.New("RoomTitle is null"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("NewRoom:"+td.title, func(t *testing.T) {
			t.Parallel()

			output1, output2 := domain.NewRoom(td.input1, td.input2)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
