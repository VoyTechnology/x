package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	pb "voytechnology/x/go/wasm-grpc/proto"
)

var (
	rpcPortFlag  = flag.Int("rpc_port", 9000, "rpc port")
	httpPortFlag = flag.Int("http_port", 8080, "http port")
	dirFlag      = flag.String("dir", "web", "directory to serve")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *rpcPortFlag))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServiceServer(s, &server{})

	go func() {
		log.Println("starting grpc server")
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	log.Println("starting http server")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *httpPortFlag), http.FileServer(http.Dir(*dirFlag))); err != nil {
		panic(err)
	}
}

type server struct{}

func (s *server) Ping(_ context.Context, req *pb.Message) (*pb.Message, error) {
	return req, nil
}
