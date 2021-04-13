package message

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	messageDomain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	mysql "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/message"

	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdCommonString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/common"
	tdMessageString "github.com/karamaru-alpha/chat-go-server/test/testdata/string/message"
)

// TestToDTO メッセージエンティティをDB情報を持った構造体に変換する処理のテスト
func TestToDTO(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		input    messageDomain.Message
		expected mysql.Message
	}{
		{
			title: "【正常系】メッセージエンティティをDB情報を持った構造体に変換",
			input: tdDomain.Entity,
			expected: mysql.Message{
				ID:     tdCommonString.ULID.Valid,
				RoomID: tdCommonString.ULID.Valid,
				Body:   tdMessageString.Body.Valid,
			},
		},
		{
			title:    "【正常系】エンティティのゼロ値が来たらからdtoを返す",
			input:    messageDomain.Message{},
			expected: mysql.Message{},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("ToDTO:"+td.title, func(t *testing.T) {
			t.Parallel()

			output := mysql.ToDTO(td.input)

			assert.Equal(t, td.expected, output)
		})
	}
}

// TestToEntity DB情報を持ったトークルームDTOをEntityに変換する処理のテスト
func TestToEntity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title     string
		input     mysql.Message
		expected1 messageDomain.Message
		expected2 error
	}{
		{
			title: "【正常系】DB情報を持った構造体をメッセージエンティティに変換",
			input: mysql.Message{
				ID:     tdCommonString.ULID.Valid,
				RoomID: tdCommonString.ULID.Valid,
				Body:   tdMessageString.Body.Valid,
			},
			expected1: tdDomain.Entity,
			expected2: nil,
		},
		{
			title:     "【正常系】'空のdtoが来たらエンティティのゼロ値を返す",
			input:     mysql.Message{},
			expected1: messageDomain.Message{},
			expected2: nil,
		},
		{
			title: "【異常系】IDが不正値",
			input: mysql.Message{
				ID:     tdCommonString.ULID.Invalid,
				RoomID: tdCommonString.ULID.Valid,
				Body:   tdMessageString.Body.Valid,
			},
			expected1: messageDomain.Message{},
			expected2: errors.New("ulid: bad data size when unmarshaling"),
		},
		{
			title: "【異常系】RoomIDが不正値",
			input: mysql.Message{
				ID:     tdCommonString.ULID.Valid,
				RoomID: tdCommonString.ULID.Invalid,
				Body:   tdMessageString.Body.Valid,
			},
			expected1: messageDomain.Message{},
			expected2: errors.New("ulid: bad data size when unmarshaling"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("ToEntity:"+td.title, func(t *testing.T) {
			t.Parallel()

			output1, output2 := mysql.ToEntity(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
