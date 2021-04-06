package message

import (
	"testing"

	"github.com/stretchr/testify/assert"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
)

// TestCreate メッセージの生成処理を担うファクトリのテスト
func TestCreate(t *testing.T) {
	t.Parallel()

	factory := domainModel.NewFactory(mockUtil.NewULIDGenerator())

	tests := []struct {
		body      string
		input1    *domainModel.RoomID
		input2    *domainModel.Body
		expected1 *domainModel.Message
		expected2 error
	}{
		{
			body:      "【正常系】",
			input1:    &tdDomain.Message.RoomID,
			input2:    &tdDomain.Message.Body,
			expected1: &tdDomain.Message.Entity,
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
