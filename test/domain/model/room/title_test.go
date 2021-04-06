package room

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

// TestNewTitle トークルーム名値オブジェクトコンストラクタのテスト
func TestNewTitle(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		input     string
		expected1 *domainModel.Title
		expected2 error
	}{
		{
			title:     "【正常系】",
			input:     tdString.Room.Title.Valid,
			expected1: &tdDomain.Room.Title,
			expected2: nil,
		},
		{
			title:     "【異常系】タイトルが空",
			input:     "",
			expected1: nil,
			expected2: errors.New("RoomTitle is null"),
		},
		{
			title:     "【異常系】タイトルが短い",
			input:     tdString.Room.Title.TooShort,
			expected1: nil,
			expected2: errors.New("RoomTitle should be Three to twenty characters"),
		},
		{
			title:     "【異常系】タイトルが長い",
			input:     tdString.Room.Title.TooLong,
			expected1: nil,
			expected2: errors.New("RoomTitle should be Three to twenty characters"),
		},
	}

	for _, td := range tests {
		td := td
		t.Run("NewTitle:"+td.title, func(t *testing.T) {

			output1, output2 := domainModel.NewTitle(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
