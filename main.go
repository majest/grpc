package main

import (
	"fmt"
	"github.com/majest/grpc/handler"
	"github.com/majest/grpc/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", ":50153") // create the listener
	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	s := handler.Server{} // Instantiate the chat server handler
	fmt.Println("Starting server")
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s) // Register chat

	if err := grpcServer.Serve(lis); err != nil { /// RUn the server
		log.Fatalf("Err2: %v", err)
	}
}
