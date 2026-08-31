package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ayushmehta03/grpc-golang/proto"
)

func callSayHelloClientStream(client  pb.GreetServiceClient, names *pb.NamesList){
	log.Printf("client streaming started")
   stream, err := client.SayHelloClientStreaming(context.Background())
   if err!=nil{
	log.Fatalf("could not send naes %v",err)
   }
   for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		log.Printf("Sent request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("%v", res.Messages)



}