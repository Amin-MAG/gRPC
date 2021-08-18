package main

import (
	"context"
	"echo/api"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

const address = "localhost:8888"

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("cannot connect:", err)
	}
	defer func() {
		_ = conn.Close()
	}()

	// Setup the stream
	c := api.NewEchoClient(conn)
	stream, err := c.GlobalCounter(context.Background(), &api.CounterParam{})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("closing the stream from client")
			break
		}
		if err != nil {
			fmt.Println("Failed to receive a note :", err)
			break
		}

		fmt.Println("Counter is showing", in.Counter)
	}

	_ = stream.CloseSend()
	time.Sleep(2 * time.Second)
}
