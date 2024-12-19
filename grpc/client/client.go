package main

import (
	"context"
	"grpc/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Shomi Khan", Email: "nkskl6@gmail.com"})
	if err != nil {
		log.Fatalf("Error while calling SayHello: %v", err)
	}

	log.Printf("Response from server: %s", resp.GetMessage())
}
