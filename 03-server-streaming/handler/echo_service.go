package handler

import (
	"echo/api"
	"echo/player"
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
	//go count()

	return &echoService{
		Server: s,
	}
}

func (ps *echoService) GlobalRadio(cp *api.RadioParams, stream api.Echo_GlobalRadioServer) error {
	log.Println("Entering Global Radio")
	for {
		// Get bytes from file
		data, err := player.GetTestFileChunks()
		if err != nil {
			return err
		}

		for i := 0; i < len(data); i++ {
			err = stream.Send(&data[i])
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func (ps *echoService) GlobalCounter(cp *api.CounterParam, stream api.Echo_GlobalCounterServer) error {
	log.Println("Entering Global Counter")
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
