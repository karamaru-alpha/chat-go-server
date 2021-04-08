package message_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	dto "github.com/karamaru-alpha/chat-go-server/interfaces/dto/message"
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/message"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

// TestToProto EntityをgRPC型のモデルに変換する関数のテスト
func TestToProto(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		input    *domain.Message
		expected *pb.Message
	}{
		{
			title: "【正常系】メッセージエンティティをDB情報を持った構造体に変換する",
			input: &tdDomain.Message.Entity,
			expected: &pb.Message{
				Id:     tdString.Message.ID.Valid,
				RoomId: tdString.Room.ID.Valid,
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

		t.Run("ToProto:"+td.title, func(t *testing.T) {
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
		input    *[]domain.Message
		expected []*pb.Message
	}{
		{
			title: "【正常系】メッセージエンティティのスライスをDB情報を持った構造体に変換する",
			input: &[]domain.Message{tdDomain.Message.Entity, tdDomain.Message.Entity},
			expected: []*pb.Message{
				{
					Id:     tdString.Message.ID.Valid,
					RoomId: tdString.Room.ID.Valid,
					Body:   tdString.Message.Body.Valid,
				},
				{
					Id:     tdString.Message.ID.Valid,
					RoomId: tdString.Room.ID.Valid,
					Body:   tdString.Message.Body.Valid,
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
			output := dto.ToProtos(td.input)

			assert.Equal(t, td.expected, output)
		})
	}
}
