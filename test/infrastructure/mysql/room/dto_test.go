package room

import (
	"errors"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	infra "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/room"
	mockUtil "github.com/karamaru-alpha/chat-go-server/mock/util"
	testdata "github.com/karamaru-alpha/chat-go-server/test/testdata"
)

// TestToDTO トークルームEntityをDB情報を持つDTOに変換する処理のテスト
func TestToDTO(t *testing.T) {
	t.Parallel()

	// DTOに入れるEntityの準備
	roomTitle, err := domainModel.NewTitle(testdata.Room.Title.Valid)
	assert.NoError(t, err)

	factory := domainModel.NewFactory(mockUtil.GenerateULID)

	room, err := factory.Create(roomTitle)
	assert.NoError(t, err)

	tests := []struct {
		title    string
		input    *domainModel.Room
		expected *infra.Room
	}{
		{
			title:    "【正常系】",
			input:    room,
			expected: &infra.Room{ID: ulid.ULID(room.ID).String(), Title: string(room.Title)},
		},
	}

	for _, td := range tests {
		td := td

		output := infra.ToDTO(td.input)

		assert.Equal(t, td.expected, output)
	}
}

// TestToEntity DB情報を持ったトークルームDTOをEntityに変換する処理のテスト
func TestToEntity(t *testing.T) {
	t.Parallel()

	// Entityに変換するDTOの準備
	dto := &infra.Room{ID: testdata.Room.ID.ValidPlain, Title: testdata.Room.Title.Valid}

	// DTOから生成されるであろうEntityの準備
	roomTitle, err := domainModel.NewTitle(testdata.Room.Title.Valid)
	assert.NoError(t, err)

	factory := domainModel.NewFactory(mockUtil.GenerateULID)

	room, err := factory.Create(roomTitle)
	assert.NoError(t, err)

	tests := []struct {
		title     string
		input     *infra.Room
		expected1 *domainModel.Room
		expected2 error
	}{
		{
			title:     "【正常系】",
			input:     dto,
			expected1: room,
			expected2: nil,
		},
		{
			title:     "【異常系】タイトルが不正値(short)",
			input:     &infra.Room{ID: testdata.Room.ID.ValidPlain, Title: testdata.Room.Title.TooShort},
			expected1: nil,
			expected2: errors.New("RoomTitle should be Three to twenty characters"),
		},
		{
			title:     "【異常系】タイトルが不正値(long)",
			input:     &infra.Room{ID: testdata.Room.ID.ValidPlain, Title: testdata.Room.Title.TooLong},
			expected1: nil,
			expected2: errors.New("RoomTitle should be Three to twenty characters"),
		},
	}

	for _, td := range tests {
		td := td

		output1, output2 := infra.ToEntity(td.input)

		assert.Equal(t, td.expected1, output1)
		assert.Equal(t, td.expected2, output2)
	}
}

// TestToEntities DB情報を持ったトークルームDTOをEntityのリストに変換する処理のテスト
func TestToEntities(t *testing.T) {
	t.Parallel()

	// Entityに変換するDTOの準備
	dto := infra.Room{ID: testdata.Room.ID.ValidPlain, Title: testdata.Room.Title.Valid}

	// DTOから生成されるであろうEntityの準備
	roomTitle, err := domainModel.NewTitle(testdata.Room.Title.Valid)
	assert.NoError(t, err)

	factory := domainModel.NewFactory(mockUtil.GenerateULID)

	room, err := factory.Create(roomTitle)
	assert.NoError(t, err)

	tests := []struct {
		title     string
		input     *[]infra.Room
		expected1 *[]domainModel.Room
		expected2 error
	}{
		{
			title:     "【正常系】1つのDTOをEntityに変換",
			input:     &[]infra.Room{dto},
			expected1: &[]domainModel.Room{*room},
			expected2: nil,
		},
		{
			title:     "【正常系】2つのDTOをEntityに変換",
			input:     &[]infra.Room{dto, dto},
			expected1: &[]domainModel.Room{*room, *room},
			expected2: nil,
		},
		{
			title:     "【異常系】タイトルが不正値(short)",
			input:     &[]infra.Room{{ID: testdata.Room.ID.ValidPlain, Title: testdata.Room.Title.TooShort}},
			expected1: nil,
			expected2: errors.New("RoomTitle should be Three to twenty characters"),
		},
		{
			title:     "【異常系】タイトルが不正値(long)",
			input:     &[]infra.Room{{ID: testdata.Room.ID.ValidPlain, Title: testdata.Room.Title.TooLong}},
			expected1: nil,
			expected2: errors.New("RoomTitle should be Three to twenty characters"),
		},
	}

	for _, td := range tests {
		td := td

		output1, output2 := infra.ToEntities(td.input)

		assert.Equal(t, td.expected1, output1)
		assert.Equal(t, td.expected2, output2)
	}
}
