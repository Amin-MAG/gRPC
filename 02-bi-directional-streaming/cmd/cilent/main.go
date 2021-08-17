package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"prime/api"
	"strconv"
	"strings"
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
	c := api.NewFactorsClient(conn)
	stream, err := c.CalculatePrimeFactors(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	wc := make(chan struct{})
	go receiveFromServer(wc, stream)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a number: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "q" {
			break
		}

		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("bad input")
		}

		_ = stream.Send(&api.CalculatePrimeFactorsRequest{
			// Input new number
			Num: int64(num),
		})

		<-wc
	}

	_ = stream.CloseSend()
	time.Sleep(2 * time.Second)
}

func receiveFromServer(wc chan struct{}, stream api.Factors_CalculatePrimeFactorsClient) {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("closing the stream from client")
			return
		}
		if err != nil {
			fmt.Println("Failed to receive a note :", err)
		}

		if in.Done {
			wc <- struct{}{}
		} else {
			fmt.Println("Got new factor:", in.Result)
		}
	}
}
