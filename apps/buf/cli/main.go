package main

import (
	"context"
	"log"
	"time"

	pb "github.com/webteleport/ufo/apps/buf/greet/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address     = "buf.ufo.k0s.io:443"
	defaultName = "world"
)

func main() {
	// Set up a secure connection to the server.
	creds := credentials.NewClientTLSFromCert(nil, "")
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Greet(ctx, &pb.GreetRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetGreeting())
}
