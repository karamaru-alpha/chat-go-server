package message

import (
	"errors"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

// TestNewRoomID メッセージが紐づくトークルームの識別子を表す値オブジェクトコンストラクタのテスト
func TestNewRoomID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		input     *ulid.ULID
		expected1 *domainModel.RoomID
		expected2 error
	}{
		{
			title:     "【正常系】",
			input:     &tdULID.Room.ID,
			expected1: &tdDomain.Message.RoomID,
			expected2: nil,
		},
		{
			title:     "【異常系】引数がnil",
			input:     nil,
			expected1: nil,
			expected2: errors.New("MessageRoomID is null"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("NewRoomID:"+td.title, func(t *testing.T) {

			output1, output2 := domainModel.NewRoomID(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
