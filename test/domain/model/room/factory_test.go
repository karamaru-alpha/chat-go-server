package room

import (
	"testing"

	"github.com/stretchr/testify/assert"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	testdata "github.com/karamaru-alpha/chat-go-server/test/testdata"
)

// TestCreate トークルーム生成処理を担うファクトリのテスト
func TestCreate(t *testing.T) {
	t.Parallel()

	roomTitle, err := domainModel.NewTitle(testdata.Room.Title.Valid)
	assert.NoError(t, err)

	tests := []struct {
		title      string
		input      *domainModel.Title
		expected1T interface{} // uuidのモックが大変なので型のみで判定
		expected2  error
	}{
		{
			title:      "【正常系】",
			input:      roomTitle,
			expected1T: new(domainModel.Room),
			expected2:  nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Create:"+td.title, func(t *testing.T) {

			output1, output2 := domainModel.Create(td.input)

			assert.IsType(t, td.expected1T, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
