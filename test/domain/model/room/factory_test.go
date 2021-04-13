package room

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"

	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

type testHandler struct {
	factory domain.IFactory

	ulidGenerator *mockUtil.MockIULIDGenerator
}

// TestCreate トークルーム生成処理を担うファクトリのテスト
func TestCreate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		before    func(testHandler)
		input     domain.Title
		expected1 domain.Room
		expected2 error
	}{
		{
			title: "【正常系】トークルームエンティティ生成",
			before: func(h testHandler) {
				h.ulidGenerator.EXPECT().Generate().Return(tdULID.ULID)
			},
			input:     tdDomain.Title,
			expected1: tdDomain.Entity,
			expected2: nil,
		},
		{
			title: "【異常系】Titleが空文字列",
			before: func(h testHandler) {
				h.ulidGenerator.EXPECT().Generate().Return(tdULID.ULID)
			},
			input:     "",
			expected1: domain.Room{},
			expected2: errors.New("RoomTitle is null"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Create:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)

			td.before(tester)

			output1, output2 := tester.factory.Create(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}

func (h *testHandler) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	h.ulidGenerator = mockUtil.NewMockIULIDGenerator(ctrl)

	h.factory = domain.NewFactory(h.ulidGenerator)
}
