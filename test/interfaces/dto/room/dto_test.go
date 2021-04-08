package room_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	dto "github.com/karamaru-alpha/chat-go-server/interfaces/dto/room"
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

// TestToProto EntityをgRPC型のモデルに変換する関数のテスト
func TestToProto(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		input    *domain.Room
		expected *pb.Room
	}{
		{
			title: "【正常系】トークルームエンティティをDB情報を持った構造体に変換する",
			input: &tdDomain.Room.Entity,
			expected: &pb.Room{
				Id:    tdString.Room.ID.Valid,
				Title: tdString.Room.Title.Valid,
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

		t.Run("ToProto:"+td.title, func(t *testing.T) {
			t.Parallel()

			output := dto.ToProto(td.input)

			assert.Equal(t, td.expected, output)
		})
	}
}

// TestToProtos EntityのスライスをgRPC型のモデルに変換する関数のテスト
func TestToProtos(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		input    *[]domain.Room
		expected []*pb.Room
	}{
		{
			title: "【正常系】トークルームエンティティのスライスをDB情報を持った構造体に変換する",
			input: &[]domain.Room{tdDomain.Room.Entity, tdDomain.Room.Entity},
			expected: []*pb.Room{
				{
					Id:    tdString.Room.ID.Valid,
					Title: tdString.Room.Title.Valid,
				},
				{
					Id:    tdString.Room.ID.Valid,
					Title: tdString.Room.Title.Valid,
				},
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

		t.Run("ToProtos:"+td.title, func(t *testing.T) {
			t.Parallel()

			output := dto.ToProtos(td.input)

			assert.Equal(t, td.expected, output)
		})
	}
}
