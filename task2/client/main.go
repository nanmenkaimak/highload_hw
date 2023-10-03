package main

import (
	"context"
	pb "github.com/nanmenkaimak/highload_hw/task2/protos/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const portNumber = ":8080"

var messageBody = "hello"

func main() {
	conn, err := grpc.Dial(portNumber, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err.Error())
	}
	defer conn.Close()
	c := pb.NewMessageClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SendMessage(ctx, &pb.MessageRequest{Message: messageBody})
	if err != nil {
		log.Fatalf("could not get: %s", err.Error())
	}
	log.Printf("Get message: %s", r.GetMessage())
}
