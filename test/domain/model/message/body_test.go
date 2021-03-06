package message

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/message"

	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/message"
)

// TestNewBody メッセージ本文の値オブジェクトコンストラクタのテスト
func TestNewBody(t *testing.T) {
	t.Parallel()

	tests := []struct {
		body      string
		input     string
		expected1 domain.Body
		expected2 error
	}{
		{
			body:      "【正常系】メッセージ本文値オブジェクト生成",
			input:     tdString.Body.Valid,
			expected1: tdDomain.Body,
			expected2: nil,
		},
		{
			body:      "【異常系】本文が空文字列",
			input:     tdString.Body.Empty,
			expected1: "",
			expected2: errors.New("MessageBody is empty"),
		},
		{
			body:      "【異常系】本文が長い",
			input:     tdString.Body.TooLong,
			expected1: "",
			expected2: errors.New("MessageBody should be 1 to 255 characters"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("NewBody:"+td.body, func(t *testing.T) {
			t.Parallel()

			output1, output2 := domain.NewBody(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
