package handler

import (
	"echo/api"
	"google.golang.org/grpc"
	"log"
	"time"
)

type echoService struct {
	*grpc.Server
	api.UnimplementedEchoServer
}

var globalCounter int64 = 0

func NewEchoService(s *grpc.Server) *echoService {
	go count()

	return &echoService{
		Server: s,
	}
}

func (ps *echoService) GlobalCounter(cp *api.CounterParam, stream api.Echo_GlobalCounterServer) error {
	log.Println("Entering Echo Player")
	for {
		time.Sleep(250 * time.Millisecond)
		resp := &api.GlobalCounterResponse{
			Counter: globalCounter,
		}
		err := stream.Send(resp)
		if err != nil {
			return err
		}
	}
}

func count() {
	for {
		globalCounter++
		time.Sleep(500 * time.Millisecond)
		log.Println(globalCounter)
	}
}
