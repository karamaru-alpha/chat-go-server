package room

import (
	"testing"

	"github.com/stretchr/testify/assert"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
)

// TestCreate トークルーム生成処理を担うファクトリのテスト
func TestCreate(t *testing.T) {
	t.Parallel()

	factory := domain.NewFactory(mockUtil.NewULIDGenerator())

	tests := []struct {
		title     string
		input     *domain.Title
		expected1 *domain.Room
		expected2 error
	}{
		{
			title:     "【正常系】",
			input:     &tdDomain.Room.Title,
			expected1: &tdDomain.Room.Entity,
			expected2: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Create:"+td.title, func(t *testing.T) {

			output1, output2 := factory.Create(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
