package main

import (
	"echo/api"
	"echo/handler"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = 8888

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("ERROR:", err.Error())
	}

	grpcServer := grpc.NewServer()
	echoService := handler.NewEchoService(grpcServer)

	api.RegisterEchoServer(grpcServer, echoService)

	log.Println("Starting Server on port", port)
	_ = grpcServer.Serve(lis)
}
