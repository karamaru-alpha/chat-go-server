package room_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	domainModel "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	controller "github.com/karamaru-alpha/chat-go-server/interfaces/controller/room"
	pb "github.com/karamaru-alpha/chat-go-server/interfaces/proto/pb"
	tdDomain "github.com/karamaru-alpha/chat-go-server/test/testdata/domain/room"
	tdString "github.com/karamaru-alpha/chat-go-server/test/testdata/string"
)

// TestToProto EntityをgRPC型のモデルに変換する関数のテスト
func TestToProto(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		input    *domainModel.Room
		expected *pb.Room
	}{
		{
			title: "【正常系】",
			input: &tdDomain.Room.Entity,
			expected: &pb.Room{
				Id:    tdString.Room.ID.Valid,
				Title: tdString.Room.Title.Valid,
			},
		},
	}

	for _, td := range tests {
		td := td

		output := controller.ToProto(td.input)

		assert.Equal(t, td.expected, output)
	}
}

// TestToProtos EntityのスライスをgRPC型のモデルに変換する関数のテスト
func TestToProtos(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		input    *[]domainModel.Room
		expected []*pb.Room
	}{
		{
			title: "【正常系】",
			input: &[]domainModel.Room{tdDomain.Room.Entity, tdDomain.Room.Entity},
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
	}

	for _, td := range tests {
		td := td

		output := controller.ToProtos(td.input)

		assert.Equal(t, td.expected, output)
	}
}
