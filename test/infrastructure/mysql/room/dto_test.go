package room

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	infra "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/room"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

// TestToDTO トークルームEntityをDB情報を持つDTOに変換する処理のテスト
func TestToDTO(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		input    *domain.Room
		expected *infra.Room
	}{
		{
			title:    "【正常系】メッセージエンティティをDB情報を持った構造体に変換",
			input:    &tdDomain.Room.Entity,
			expected: &infra.Room{ID: tdString.Room.ID.Valid, Title: tdString.Room.Title.Valid},
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
		input     *infra.Room
		expected1 *domain.Room
		expected2 error
	}{
		{
			title:     "【正常系】DB情報を持った構造体をトークルームエンティティに変換",
			input:     &infra.Room{ID: tdString.Room.ID.Valid, Title: tdString.Room.Title.Valid},
			expected1: &tdDomain.Room.Entity,
			expected2: nil,
		},
		{
			title:     "【正常系】nilが来たらnilを返す",
			input:     nil,
			expected1: nil,
			expected2: nil,
		},
		{
			title:     "【異常系】IDが不正値",
			input:     &infra.Room{ID: tdString.Room.ID.Invalid, Title: tdString.Room.Title.Valid},
			expected1: nil,
			expected2: errors.New("ulid: bad data size when unmarshaling"),
		},
		{
			title:     "【異常系】タイトルが空文字列",
			input:     &infra.Room{ID: tdString.Room.ID.Valid, Title: ""},
			expected1: nil,
			expected2: errors.New("RoomTitle is null"),
		},
		{
			title:     "【異常系】タイトルが不正値(short)",
			input:     &infra.Room{ID: tdString.Room.ID.Valid, Title: tdString.Room.Title.TooShort},
			expected1: nil,
			expected2: errors.New("RoomTitle should be Three to twenty characters"),
		},
		{
			title:     "【異常系】タイトルが不正値(long)",
			input:     &infra.Room{ID: tdString.Room.ID.Valid, Title: tdString.Room.Title.TooLong},
			expected1: nil,
			expected2: errors.New("RoomTitle should be Three to twenty characters"),
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

// TestToEntities DB情報を持ったトークルームDTOをEntityのリストに変換する処理のテスト
func TestToEntities(t *testing.T) {
	t.Parallel()

	// Entityに変換するDTOの準備
	dto := infra.Room{ID: tdString.Room.ID.Valid, Title: tdString.Room.Title.Valid}

	tests := []struct {
		title     string
		input     *[]infra.Room
		expected1 *[]domain.Room
		expected2 error
	}{
		{
			title:     "【正常系】1つのDTOをEntityに変換",
			input:     &[]infra.Room{dto},
			expected1: &[]domain.Room{tdDomain.Room.Entity},
			expected2: nil,
		},
		{
			title:     "【正常系】2つのDTOをEntityに変換",
			input:     &[]infra.Room{dto, dto},
			expected1: &[]domain.Room{tdDomain.Room.Entity, tdDomain.Room.Entity},
			expected2: nil,
		},
		{
			title:     "【正常系】nilが来たらnilを返す",
			input:     nil,
			expected1: nil,
			expected2: nil,
		},
		{
			title:     "【異常系】タイトルが不正値(short)",
			input:     &[]infra.Room{{ID: tdString.Room.ID.Valid, Title: tdString.Room.Title.TooShort}},
			expected1: nil,
			expected2: errors.New("RoomTitle should be Three to twenty characters"),
		},
		{
			title:     "【異常系】タイトルが不正値(long)",
			input:     &[]infra.Room{{ID: tdString.Room.ID.Valid, Title: tdString.Room.Title.TooLong}},
			expected1: nil,
			expected2: errors.New("RoomTitle should be Three to twenty characters"),
		},
	}

	for _, td := range tests {
		td := td

		t.Run("ToEntities:"+td.title, func(t *testing.T) {
			output1, output2 := infra.ToEntities(td.input)

			assert.Equal(t, td.expected1, output1)
			assert.Equal(t, td.expected2, output2)
		})
	}
}
