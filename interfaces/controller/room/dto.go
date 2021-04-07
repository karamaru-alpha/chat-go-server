package room

import (
	"github.com/oklog/ulid"

	domain "github.com/karamaru-alpha/chat-go-server/domain/model/room"
	pb "github.com/karamaru-alpha/chat-go-server/interfaces/proto/pb"
)

// ToProto トークルームエンティティをｇRPCの型に変換
func ToProto(entity *domain.Room) *pb.Room {
	return &pb.Room{
		Id:    ulid.ULID(entity.ID).String(),
		Title: string(entity.Title),
	}
}

// ToProtos トークルームエンティティのスライスをｇRPCの型に変換
func ToProtos(entities *[]domain.Room) []*pb.Room {
	var rooms []*pb.Room
	for _, v := range *entities {
		rooms = append(rooms, ToProto(&v))
	}
	return rooms
}
