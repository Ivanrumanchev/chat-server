package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/Ivanrumanchev/chat-server/pkg/chat_v1"
)

func generateChatID() int64 {
	t := time.Now().UnixNano() / int64(time.Millisecond)
	return t
}

const grpcPort = 50151

type server struct {
	desc.UnimplementedChatV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Create
func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create UserIDs: %s", req.GetUserIDs())
	log.Printf("Create Name: %s", req.GetName())

	return &desc.CreateResponse{
		Id: generateChatID(),
	}, nil
}

// Delete
func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete Chat id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}

// SendMessage
func (s *server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("SendMessage From: %s", req.GetFrom())
	log.Printf("SendMessage Text: %s", req.GetText())
	log.Printf("SendMessage Timestamp: %s", req.GetTimestamp())

	return &emptypb.Empty{}, nil
}
