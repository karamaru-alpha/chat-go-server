package message

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
)

// TestNewMessage メッセージエンティティコンストラクタのテスト
func TestNewMessage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		body      string
		input1    *domainModel.ID
		input2    *domainModel.RoomID
		input3    *domainModel.Body
		expected1 *domainModel.Message
		expected2 error
	}{
		{
			body:      "【正常系】",
			input1:    &tdDomain.Message.ID,
			input2:    &tdDomain.Message.RoomID,
			input3:    &tdDomain.Message.Body,
			expected1: &tdDomain.Message.Entity,
			expected2: nil,
		},
		{
			body:      "【異常系】IDがnil",
			input1:    nil,
			input2:    &tdDomain.Message.RoomID,
			input3:    &tdDomain.Message.Body,
			expected1: nil,
			expected2: errors.New("MessageID is null"),
		},
		{
			body:      "【異常系】RoomIDがnil",
			input1:    &tdDomain.Message.ID,
			input2:    nil,
			input3:    &tdDomain.Message.Body,
			expected1: nil,
			expected2: errors.New("MessageRoomID is null"),
		},
		{
			body:      "【異常系】Bodyがnil",
			input1:    &tdDomain.Message.ID,
			input2:    &tdDomain.Message.RoomID,
			input3:    nil,
			expected1: nil,
			expected2: errors.New("MessageBody is null"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("NewMessage:"+td.body, func(t *testing.T) {

			output1, output2 := domainModel.NewMessage(td.input1, td.input2, td.input3)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
