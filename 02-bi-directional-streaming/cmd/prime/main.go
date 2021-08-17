package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"prime/api"
	"prime/handler"
)

const port = 8888

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("ERROR:", err.Error())
	}

	grpcServer := grpc.NewServer()
	primeService := handler.NewPrimeService(grpcServer)

	api.RegisterFactorsServer(grpcServer, primeService)

	log.Println("Starting Server on port", port)
	_ = grpcServer.Serve(lis)
}
