package message

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	roomDomain "github.com/karamaru-alpha/chat-go-server/domain/model/room"

	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	tdMessageDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdRoomDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdULID "github.com/karamaru-alpha/chat-go-server/test/testdata/ulid"
)

type testHandler struct {
	factory messageDomain.IFactory

	ulidGenerator *mockUtil.MockIULIDGenerator
}

// TestCreate メッセージの生成処理を担うファクトリのテスト
func TestCreate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		before    func(testHandler)
		input1    roomDomain.Room
		input2    messageDomain.Body
		expected1 messageDomain.Message
		expected2 error
	}{
		{
			title: "【正常系】メッセージエンティティ生成",
			before: func(h testHandler) {
				h.ulidGenerator.EXPECT().Generate().Return(tdULID.ULID)
			},
			input1:    tdRoomDomain.Entity,
			input2:    tdMessageDomain.Body,
			expected1: tdMessageDomain.Entity,
			expected2: nil,
		},
		{
			title:     "【異常系】Roomがゼロ値",
			input1:    roomDomain.Room{},
			input2:    tdMessageDomain.Body,
			expected1: messageDomain.Message{},
			expected2: errors.New("MessageRoom is not exist"),
		},
		{
			title: "【異常系】Bodyがゼロ値",
			before: func(h testHandler) {
				h.ulidGenerator.EXPECT().Generate().Return(tdULID.ULID)
			},
			input1:    tdRoomDomain.Entity,
			input2:    "",
			expected1: messageDomain.Message{},
			expected2: errors.New("MessageBody is null"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("Create:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)

			if td.before != nil {
				td.before(tester)
			}

			output1, output2 := tester.factory.Create(td.input1, td.input2)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}

func (h *testHandler) setupTest(t *testing.T) {
	ctrl := gomock.NewController(t)
	h.ulidGenerator = mockUtil.NewMockIULIDGenerator(ctrl)

	h.factory = messageDomain.NewFactory(h.ulidGenerator)
}
