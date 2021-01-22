package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "hello/proto"
	"log"
	"net"
)

const (
	port    = ":50051"
	network = "127.0.0.1"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	fmt.Println("vim-go")
	lis, err := net.Listen("tcp", port)
	fmt.Println("net.Listen", lis.Addr().String())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
