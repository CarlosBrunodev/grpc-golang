package main

import (
	"context"
	"log"

	"github.com/CarlosBrunodev/grpc-golang/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main(){
	conn, err := grpc.Dial("localhost:9000",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewHelloClient(conn)

	req := &pb.HelloRequest{
		Name: "John Doe",
	}

	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
}