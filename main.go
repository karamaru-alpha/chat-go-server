package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	roomDI "github.com/karamaru-alpha/chat-go-server/di/room"
	pb "github.com/karamaru-alpha/chat-go-server/proto/pb"
)

func main() {
	port := os.Getenv("PORT")

	// リスナー登録
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Could not listen @ %v :: %v", port, err)
	}
	log.Printf("Listening @ :%s", port)

	// gRPC作成
	grpcServer := grpc.NewServer()
	pb.RegisterRoomServicesServer(grpcServer, roomDI.DI())

	// サーバ起動
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to start gRPC Server :: %v", err)
	}
}
