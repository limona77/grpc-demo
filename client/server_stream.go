package main

import (
	"context"
	pb "grpc-demo/proto"
	"io"
	"log"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Print("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming: %v", err)
		}
		log.Println(message)
	}
	log.Print("Streaming finished")
}
