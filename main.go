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

	lis, err := net.Listen("tcp", ":50153") // Tworzenie nowego ucha na porcie 50153 na protokole TCP
	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	s := handler.Server{} // Tworzernie serwera
	fmt.Println("Starting server")
	grpcServer := grpc.NewServer() // Tworzenie serwera GRPC ( Jeszcze nie wiem czym się różnią )

	chat.RegisterChatServiceServer(grpcServer, &s) // Rejestrowanie nowego klienta chat serwisu

	if err := grpcServer.Serve(lis); err != nil { // Usatwiamy By serwer GRPC korzystał z net.Listener'a przy okazji prawdzamy czy się udało. ( coś jak if(listener.listen(50153) != sf::Socket::Done) { std::cout << "Error" << std::endl; } ) )
		log.Fatalf("Err2: %v", err)
	}

	fmt.Println("2") // Linijka do debugu

}
