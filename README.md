# GRPC
Pierwszy kod + wskazówki

<h1> Kod przerżnięty z kursu https://tutorialedge.net/golang/go-grpc-beginners-tutorial/ </h1> </br>
Coś ogarnąłem, ale chyba nie wystarczająco.  </br> </br>

Teraz wiem jak: kompilować za pomocą protoc, stworzyć prosty serwer, stworzyć prostego kleinta.  </br> </br>

Dalej nie wiem: po co mi "github.com/tutorialedge/go-grpc-tutorial/chat" - coś importuję, ale co? </br>
Czego kontekst bierze "context" / "context.Background()"; </br>
Za co odpowiadają pliki: chat.pb.go i chat_grpc.pb.go.
  
  <h2> Wskazówki </h2>
  Jak kompilowac proto: protoc -I $GOPATH/src/ --go_out . --go_opt paths=source_relative --go-grpc_out . --go-grpc_opt paths=source_relative /Users/nintyswinty/go/src/Zabawa_z_gRPC/pb/chat.proto </br> </br>
  Nie uruchamiać go z terminala tylko z GoLand'a. ( Przeczytałęm że go run żadko wychodzi - lepiej używać go build + w GoLand wykrywa mi GoPath ).  </br> </br>
  Robić folder pb.  </br> </br>
  
  <h3> Błędy </h3> </br>
  
  Nie wiem dlaczego ale kiedy serwer wysyła do client'a widomośc to otrzymuję "Hello From the Server!" gdzie w chat.go jest "Hello from the server!". Podjerzewam że to wina tego gita w importach.
  
  
