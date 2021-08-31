package main

import (
	"context"
	"log"

	pingpong "github.com/bogdan-user/go-grpc-react/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	
	certFile := "../cert/ca.crt" // Certificate Authority Trust certificate
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")

	if sslErr != nil {
			log.Fatalf("Error while loading SSL certificate: %v\n", sslErr)
			
	}
	opts := grpc.WithTransportCredentials(creds)
	
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect to: %v\n", err)
	}
	defer cc.Close()

	c := pingpong.NewPingPongClient(cc)

	pong, err := c.Ping(context.Background(), &pingpong.PingRequest{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(pong)
}
