package message

import (
	"errors"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/message"

	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

// TestNewID メッセージ識別子を表す値オブジェクトコンストラクタのテスト
func TestNewID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		input     ulid.ULID
		expected1 domain.ID
		expected2 error
	}{
		{
			title:     "【正常系】メッセージの識別子生成",
			input:     tdULID.ULID,
			expected1: tdDomain.ID,
			expected2: nil,
		},
		{
			title:     "【異常系】ulidがゼロ値",
			input:     ulid.ULID{},
			expected1: domain.ID{},
			expected2: errors.New("MessageID is null"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("NewID:"+td.title, func(t *testing.T) {
			t.Parallel()

			output1, output2 := domain.NewID(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
