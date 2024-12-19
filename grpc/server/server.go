package main

import (
	"context"
	"fmt"
	"grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server implementation
type HelloServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	str := fmt.Sprintf("name: %s, email: %s", req.GetName(), req.GetEmail())
	log.Println("Received:", str)
	return &pb.HelloResponse{Message: str}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &HelloServer{})

	log.Println("Server is running on port 50050...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
