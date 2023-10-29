package main

import (
	"log"
	"net"

	"github.com/akshayUr04/totalitycorp-service/pkg/config"
	"github.com/akshayUr04/totalitycorp-service/pkg/pb"
	"github.com/akshayUr04/totalitycorp-service/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("can't load config %s", err)
	}

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("can't create listener %s", err)
	}
	log.Printf("server listening at %v", cfg.Port)

	s := service.UserServer{}
	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
