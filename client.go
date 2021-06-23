package main

import (
	"log" // Odpowiada za logi

	"github.com/tutorialedge/go-grpc-tutorial/chat" // Importuje coś z git'a - nie wiem co
	"golang.org/x/net/context" // Importuje nam context
	"google.golang.org/grpc" // Importuje GRPC
)

func main () { // Główna pętla programu
	var conn *grpc.ClientConn // Zmienna przechowująca połączenie
	conn, err := grpc.Dial(":50153", grpc.WithInsecure()); // Połączenie na porcie 50153 ( Połączenie NIEstrzeżone )
	if err != nil {
		log.Fatalf("Err: %s", err)
	}
	defer conn.Close() // Funckja wykonująca się pod wykonaniu main'a

	c := chat.NewChatServiceClient(conn) // Tworzy nowego klienta dla serwisu chat (?)

	message := chat.Message { // Ustawiamy zawartość wiadomości
		Body: "Hello from the client!",
	}

	response, err := c.SayHello(context.Background(), &message) // Używamy funkcji którą wcześniej stworzyliśmy
	if err != nil {
		log.Fatalf("Err2: %s", err)
	}

	log.Printf("Response: %s", response.Body) // Wypisujemy wiado od serwera
}
