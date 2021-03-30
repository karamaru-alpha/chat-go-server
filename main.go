package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/karamaru-alpha/chat-go-server/config"
	roomDI "github.com/karamaru-alpha/chat-go-server/di/room"
	pb "github.com/karamaru-alpha/chat-go-server/interfaces/proto/pb"
)

func main() {
	// リスナー登録
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Port()))
	if err != nil {
		log.Fatalf("Could not listen @ %v :: %v", config.Port(), err)
	}
	log.Printf("Listening @ :%s", config.Port())

	// gRPC作成
	grpcServer := grpc.NewServer()
	pb.RegisterRoomServicesServer(grpcServer, roomDI.DI())

	// サーバ起動
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}
}
