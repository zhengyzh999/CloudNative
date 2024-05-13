package main

import (
	"context"
	"flag"
	"fmt"
	"gateway/myservice/protoservice"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	protoservice.RegisterMyServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type server struct {
	protoservice.UnimplementedMyServiceServer
}

func (s *server) Echo(ctx context.Context, in *protoservice.Message) (*protoservice.Message, error) {
	fmt.Printf("server recv: %+v \n", in)
	return in, nil
}

func (s *server) EchoSimple(ctx context.Context, in *protoservice.SimpleMessage) (*protoservice.SimpleMessage, error) {
	fmt.Printf("server recv: %+v \n", in)
	return in, nil
}
