package greeter_server

import (
	"context"
	"flag"
	"fmt"
	"github.com/atadzan/simple-grpc/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

// SayHello implements helloworld.Greeter
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Helllo" + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}
	s := grpc.NewServer()

	helloworld.RegisterGreeterServer(s, &server{})
	log.Printf("server %s", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
