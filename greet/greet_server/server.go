package main

import (
	"fmt"
	"log"
	"net"

	"github.com/JyotsnaGorle/grpc-go"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Welcome")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen :%v", err)
	}

	s := grpc.NewServer()

	greetpb.Register(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
