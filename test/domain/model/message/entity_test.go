package message

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	tdMessageDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdRoomDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
)

// TestNewMessage メッセージエンティティコンストラクタのテスト
func TestNewMessage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		body      string
		input1    *messageDomain.ID
		input2    *roomDomain.ID
		input3    *messageDomain.Body
		expected1 *messageDomain.Message
		expected2 error
	}{
		{
			body:      "【正常系】メッセージエンティティ生成",
			input1:    &tdMessageDomain.Message.ID,
			input2:    &tdRoomDomain.Room.ID,
			input3:    &tdMessageDomain.Message.Body,
			expected1: &tdMessageDomain.Message.Entity,
			expected2: nil,
		},
		{
			body:      "【異常系】IDがnil",
			input1:    nil,
			input2:    &tdRoomDomain.Room.ID,
			input3:    &tdMessageDomain.Message.Body,
			expected1: nil,
			expected2: errors.New("MessageID is null"),
		},
		{
			body:      "【異常系】RoomIDがnil",
			input1:    &tdMessageDomain.Message.ID,
			input2:    nil,
			input3:    &tdMessageDomain.Message.Body,
			expected1: nil,
			expected2: errors.New("MessageRoomID is null"),
		},
		{
			body:      "【異常系】Bodyがnil",
			input1:    &tdMessageDomain.Message.ID,
			input2:    &tdRoomDomain.Room.ID,
			input3:    nil,
			expected1: nil,
			expected2: errors.New("MessageBody is null"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("NewMessage:"+td.body, func(t *testing.T) {
			t.Parallel()

			output1, output2 := messageDomain.NewMessage(td.input1, td.input2, td.input3)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
