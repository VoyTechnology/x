package main

import (
	"context"
	"time"

	"google.golang.org/grpc"

	pb "voytechnology/x/go/wasm-grpc/proto"
)

// Ping allows to ping a server
func main() {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewEchoServiceClient(conn)

	res, err := c.Ping(context.Background(), &pb.Message{
		Time: time.Now().Unix(),
	})
	if err != nil {
		panic(err)
	}

	panic(res.Time)
}
