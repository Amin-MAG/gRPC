package main

import (
	"bale_interview/apis"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

// for testing
const (
	address = "localhost:8989"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := apis.NewMajorityClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CalculateListMajority(ctx, &apis.ListMajorityRequest{Numbers: []int32{
		0, 0, 1, 1, 1, 1,
	}})
	if err != nil {
		log.Println(err.Error())
	}

	if r != nil {
		log.Printf("%s", r.GetMessage())
		log.Printf("%+v", r.GetMajorityNumber())
	}
}
