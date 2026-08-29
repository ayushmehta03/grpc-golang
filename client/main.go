package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
		pb "github.com/ayushmehta03/grpc-golang/proto"

)

const (
	port=":8080"
)


func main(){
	conn,err:=grpc.NewClient("localhost"+port,grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!=nil{
		log.Fatalf("did not connect: %v",err)
	}

	defer conn.Close()


client:=pb.NewGreetServiceClient(conn)

	names:=&pb.NamesList{
		Names:[]string{"am","ad","ayra","kamyar"},
}

	callSayHello(client)
	callSayHelloServerStream(client,names)
}