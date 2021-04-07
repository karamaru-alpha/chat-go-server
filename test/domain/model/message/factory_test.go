package message

import (
	"testing"

	"github.com/stretchr/testify/assert"

	messageDomainModel "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	tdMessageDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdRoomDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
)

// TestCreate メッセージの生成処理を担うファクトリのテスト
func TestCreate(t *testing.T) {
	t.Parallel()

	factory := messageDomainModel.NewFactory(mockUtil.NewULIDGenerator())

	tests := []struct {
		body      string
		input1    *roomDomainModel.ID
		input2    *messageDomainModel.Body
		expected1 *messageDomainModel.Message
		expected2 error
	}{
		{
			body:      "【正常系】",
			input1:    &tdRoomDomain.Room.ID,
			input2:    &tdMessageDomain.Message.Body,
			expected1: &tdMessageDomain.Message.Entity,
			expected2: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Create:"+td.body, func(t *testing.T) {

			output1, output2 := factory.Create(td.input1, td.input2)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
