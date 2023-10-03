package main

import (
	"context"
	pb "github.com/nanmenkaimak/highload_hw/task2/protos/message"
	"google.golang.org/grpc"
	"log"
	"net"
)

const portNumber = ":8080"

type server struct {
	pb.UnimplementedMessageServer
}

func main() {
	lis, err := net.Listen("tcp", portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}
	s := grpc.NewServer()
	pb.RegisterMessageServer(s, &server{})
	log.Printf("server listening at %s", lis.Addr().String())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}

func (s *server) SendMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	log.Printf("Received: %s", in.GetMessage())
	return &pb.MessageReply{
		Message: "Ei ne " + in.GetMessage(),
	}, nil
}
