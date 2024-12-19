package main

import (
	"context"
	"fmt"
	pb "grpc/pb/grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server implementation
type HelloServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %s", req.GetName())
	return &pb.HelloResponse{Message: fmt.Sprintf("Hello, %s!", req.GetName())}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &HelloServer{})

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
