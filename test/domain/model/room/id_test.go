package room

import (
	"errors"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
)

// TestNewID トークルームIDの値オブジェクトコンストラクタテスト
func TestNewID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		input     *ulid.ULID
		expected1 *domainModel.ID
		expected2 error
	}{
		{
			title: "【正常系】",
			input: &testData.id.valid,
			expected1: (func(v *ulid.ULID) *domainModel.ID {
				id := domainModel.ID(*v)
				return &id
			})(&testData.id.valid),
			expected2: nil,
		},
		{
			title:     "【異常系】引数がnil",
			input:     nil,
			expected1: nil,
			expected2: errors.New("RoomID is null"),
		},
	}

	for _, td := range tests {
		td := td
		t.Run("NewID:"+td.title, func(t *testing.T) {

			output1, output2 := domainModel.NewID(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
