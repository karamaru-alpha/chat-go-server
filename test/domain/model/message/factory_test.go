package message

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	tdMessageDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdRoomDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
)

// TestCreate メッセージの生成処理を担うファクトリのテスト
func TestCreate(t *testing.T) {
	t.Parallel()

	factory := messageDomain.NewFactory(mockUtil.NewULIDGenerator())

	tests := []struct {
		body      string
		input1    *roomDomain.Room
		input2    *messageDomain.Body
		expected1 *messageDomain.Message
		expected2 error
	}{
		{
			body:      "【正常系】メッセージエンティティ生成",
			input1:    &tdRoomDomain.Room.Entity,
			input2:    &tdMessageDomain.Message.Body,
			expected1: &tdMessageDomain.Message.Entity,
			expected2: nil,
		},
		{
			body:      "【異常系】Roomがnil",
			input1:    nil,
			input2:    &tdMessageDomain.Message.Body,
			expected1: nil,
			expected2: errors.New("MessageRoom is not exist"),
		},
		{
			body:      "【異常系】Bodyがnil",
			input1:    &tdRoomDomain.Room.Entity,
			input2:    nil,
			expected1: nil,
			expected2: errors.New("MessageBody is null"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Create:"+td.body, func(t *testing.T) {
			t.Parallel()

			output1, output2 := factory.Create(td.input1, td.input2)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
