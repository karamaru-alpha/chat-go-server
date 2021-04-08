package message

import (
	"github.com/oklog/ulid"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/message"
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"
)

// ToProto メッセージエンティティをｇRPC型に変換
func ToProto(entity *domain.Message) *pb.Message {
	if entity == nil {
		return nil
	}

	return &pb.Message{
		Id:     ulid.ULID(entity.ID).String(),
		RoomId: ulid.ULID(entity.RoomID).String(),
		Body:   string(entity.Body),
	}
}

// ToProtos メッセージエンティティのスライスをｇRPC型に変換
func ToProtos(entities *[]domain.Message) []*pb.Message {
	if entities == nil {
		return nil
	}

	messages := make([]*pb.Message, 0, len(*entities))
	for _, v := range *entities {
		messages = append(messages, ToProto(&v))
	}
	return messages
}
