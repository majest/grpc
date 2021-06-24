package handler

import (
	"fmt"
	"github.com/majest/grpc/chat"
	"log"
	"golang.org/x/net/context"
)

type Server struct { // Pusta struktura serwera
	chat.UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *chat.MessageRequest) (*chat.MessageResponse, error) { // Funckja która przyjmuję kontekst i wiadomość i zwraca wiadomość i error
	log.Printf("Messages: %v", message.Say) // Wypisywanie wiadomości
	return &chat.MessageResponse{Body: fmt.Sprintf("Chat server response. You said: %s", message.Say)}, nil // Zwracanie wiadomości i błędu ( a raczej jego brak )
}
