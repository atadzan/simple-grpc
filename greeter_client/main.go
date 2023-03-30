package main

import (
	"context"
	"flag"
	"github.com/atadzan/simple-grpc/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:5051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("didn't connect: %v", err.Error())
	}
	defer conn.Close()

	c := helloworld.NewGreeterClient(conn)

	// Contact the server and printout its response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err.Error())
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
