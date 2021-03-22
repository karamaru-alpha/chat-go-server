package room

import (
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	infra "github.com/karamaru-alpha/chat-go-server/infrastructure/mysql/room"
	testdata "github.com/karamaru-alpha/chat-go-server/test/testdata"
)

func TestToDTO(t *testing.T) {
	t.Parallel()

	// DTOに入れるEntityの準備
	roomTitle, err := domainModel.NewTitle(testdata.Room.Title.Valid)
	assert.NoError(t, err)
	room, err := domainModel.Create(roomTitle)
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
	roomID, err := domainModel.NewID(&testdata.Room.ID.Valid)
	assert.NoError(t, err)
	roomTitle, err := domainModel.NewTitle(testdata.Room.Title.Valid)
	assert.NoError(t, err)
	room, err := domainModel.NewRoom(roomID, roomTitle)
	assert.NoError(t, err)

	tests := []struct {
		title    string
		input    *infra.Room
		expected *domainModel.Room
		isError  bool
	}{
		{
			title:    "【正常系】",
			input:    dto,
			expected: room,
			isError:  false,
		},
		{
			title:    "【異常系】IDが不正値",
			input:    &infra.Room{ID: testdata.Room.ID.InvalidPlain, Title: testdata.Room.Title.Valid},
			expected: nil,
			isError:  true,
		},
		{
			title:    "【異常系】タイトルが不正値(short)",
			input:    &infra.Room{ID: testdata.Room.ID.ValidPlain, Title: testdata.Room.Title.TooShort},
			expected: nil,
			isError:  true,
		},
		{
			title:    "【異常系】タイトルが不正値(long)",
			input:    &infra.Room{ID: testdata.Room.ID.ValidPlain, Title: testdata.Room.Title.TooLong},
			expected: nil,
			isError:  true,
		},
	}

	for _, td := range tests {
		td := td

		output, err := infra.ToEntity(td.input)

		assert.Equal(t, td.expected, output)

		if td.isError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

// TestToEntities DB情報を持ったトークルームDTOをEntityのリストに変換する処理のテスト
func TestToEntities(t *testing.T) {
	t.Parallel()

	// Entityに変換するDTOの準備
	dto := infra.Room{ID: testdata.Room.ID.ValidPlain, Title: testdata.Room.Title.Valid}

	// DTOから生成されるであろうEntityの準備
	roomID, err := domainModel.NewID(&testdata.Room.ID.Valid)
	assert.NoError(t, err)

	roomTitle, err := domainModel.NewTitle(testdata.Room.Title.Valid)
	assert.NoError(t, err)

	room, err := domainModel.NewRoom(roomID, roomTitle)
	assert.NoError(t, err)

	tests := []struct {
		title    string
		input    *[]infra.Room
		expected *[]domainModel.Room
		isError  bool
	}{
		{
			title:    "【正常系】1つのDTOをEntityに変換",
			input:    &[]infra.Room{dto},
			expected: &[]domainModel.Room{*room},
			isError:  false,
		},
		{
			title:    "【正常系】2つのDTOをEntityに変換",
			input:    &[]infra.Room{dto, dto},
			expected: &[]domainModel.Room{*room, *room},
			isError:  false,
		},
		{
			title:    "【異常系】IDが不正値",
			input:    &[]infra.Room{{ID: testdata.Room.ID.InvalidPlain, Title: testdata.Room.Title.Valid}},
			expected: nil,
			isError:  true,
		},
		{
			title:    "【異常系】タイトルが不正値(short)",
			input:    &[]infra.Room{{ID: testdata.Room.ID.ValidPlain, Title: testdata.Room.Title.TooShort}},
			expected: nil,
			isError:  true,
		},
		{
			title:    "【異常系】タイトルが不正値(long)",
			input:    &[]infra.Room{{ID: testdata.Room.ID.ValidPlain, Title: testdata.Room.Title.TooLong}},
			expected: nil,
			isError:  true,
		},
	}

	for _, td := range tests {
		td := td

		output, err := infra.ToEntities(td.input)

		assert.Equal(t, td.expected, output)

		if td.isError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
