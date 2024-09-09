package main

import (
	"context"
	"log"
	"time"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/Ivanrumanchev/chat-server/pkg/chat_v1"
)

const (
	address = "localhost:50151"
	chatID  = int64(3)
	from    = "10"
	text    = "text"
	name    = "room3"
)

var userIDs = []string{"10", "11"}
var timestamp = timestamppb.New(time.Now())

func closeConn(conn *grpc.ClientConn) {
	if err := conn.Close(); err != nil {
		log.Printf("Error closing connection: %v", err)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer closeConn(conn)

	c := desc.NewChatV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createResponse, err := c.Create(ctx, &desc.CreateRequest{UserIDs: userIDs, Name: name})
	if err != nil {
		log.Fatalf("failed to create chat by id: %v", err)
	}

	log.Printf(color.RedString("Create Chat info:\n"), color.GreenString("%+v", createResponse))

	deleteResponse, err := c.Delete(ctx, &desc.DeleteRequest{Id: chatID})
	if err != nil {
		log.Fatalf("failed to delete chat by id: %v", err)
	}

	log.Printf(color.RedString("Delete Chat info:\n"), color.GreenString("%+v", deleteResponse))

	sendMessageResponse, err := c.SendMessage(ctx, &desc.SendMessageRequest{From: from, Text: text, Timestamp: timestamp})
	if err != nil {
		log.Fatalf("failed to send message chat by id: %v", err)
	}

	log.Printf(color.RedString("Chat info:\n"), color.GreenString("%+v", sendMessageResponse))
}
