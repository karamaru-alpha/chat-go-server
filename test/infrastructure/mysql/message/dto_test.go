package message

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	infra "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/message"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

// TestToDTO メッセージエンティティをDB情報を持った構造体に変換する処理のテスト
func TestToDTO(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		input    *messageDomain.Message
		expected *infra.Message
	}{
		{
			title: "【正常系】メッセージエンティティをDB情報を持った構造体に変換",
			input: &tdDomain.Message.Entity,
			expected: &infra.Message{
				ID:     tdString.Message.ID.Valid,
				RoomID: tdString.Room.ID.Valid,
				Body:   tdString.Message.Body.Valid,
			},
		},
		{
			title:    "【正常系】nilが来たらnilを返す",
			input:    nil,
			expected: nil,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("ToDTO:"+td.title, func(t *testing.T) {
			output := infra.ToDTO(td.input)

			assert.Equal(t, td.expected, output)
		})
	}
}

// TestToEntity DB情報を持ったトークルームDTOをEntityに変換する処理のテスト
func TestToEntity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		input     *infra.Message
		expected1 *messageDomain.Message
		expected2 error
	}{
		{
			title: "【正常系】DB情報を持った構造体をメッセージエンティティに変換",
			input: &infra.Message{
				ID:     tdString.Message.ID.Valid,
				RoomID: tdString.Room.ID.Valid,
				Body:   tdString.Message.Body.Valid,
			},
			expected1: &tdDomain.Message.Entity,
			expected2: nil,
		},
		{
			title:     "【正常系】nilが来たらnilを返す",
			input:     nil,
			expected1: nil,
			expected2: nil,
		},
		{
			title: "【異常系】IDが不正値",
			input: &infra.Message{
				ID:     tdString.Message.ID.Invalid,
				RoomID: tdString.Room.ID.Valid,
				Body:   tdString.Message.Body.Valid,
			},
			expected1: nil,
			expected2: errors.New("ulid: bad data size when unmarshaling"),
		},
		{
			title: "【異常系】RoomIDが不正値",
			input: &infra.Message{
				ID:     tdString.Message.ID.Valid,
				RoomID: tdString.Room.ID.Invalid,
				Body:   tdString.Message.Body.Valid,
			},
			expected1: nil,
			expected2: errors.New("ulid: bad data size when unmarshaling"),
		},
		{
			title: "【異常系】Bodyが不正値(empty)",
			input: &infra.Message{
				ID:     tdString.Message.ID.Valid,
				RoomID: tdString.Room.ID.Valid,
				Body:   tdString.Message.Body.Empty,
			},
			expected1: nil,
			expected2: errors.New("MessageBody is empty"),
		},
		{
			title: "【異常系】Bodyが不正値(long)",
			input: &infra.Message{
				ID:     tdString.Message.ID.Valid,
				RoomID: tdString.Room.ID.Valid,
				Body:   tdString.Message.Body.TooLong,
			},
			expected1: nil,
			expected2: errors.New("MessageBody should be 1 to 255 characters"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("ToEntity:"+td.title, func(t *testing.T) {
			output1, output2 := infra.ToEntity(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
