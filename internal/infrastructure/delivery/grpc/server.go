package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartServer(portGRPC string, services ...Service) {
	grpcServer := grpc.NewServer(
	//grpc.UnaryInterceptor(middleware.JWTInterceptor)
	)

	for _, service := range services {
		service.Register(grpcServer)
	}
	port := ":" + portGRPC

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("gRPC server listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type Service interface {
	Register(server *grpc.Server)
}
