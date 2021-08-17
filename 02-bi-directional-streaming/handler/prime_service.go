package handler

import (
	"google.golang.org/grpc"
	"io"
	"log"
	"math"
	"prime/api"
)

type primeService struct {
	*grpc.Server
	api.UnimplementedFactorsServer
}

func NewPrimeService(s *grpc.Server) *primeService {
	return &primeService{
		Server: s,
	}
}

func (ps *primeService) CalculatePrimeFactors(stream api.Factors_CalculatePrimeFactorsServer) error {
	log.Println("Entering CalculatePrimeFactors")

	for {
		log.Println("waiting for new receive")
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Println("Received:", req.Num)

		// time.Sleep(time.Second)

		c := make(chan int64)
		go findPrimeFactors(c, req.Num)
		for n := range c {
			resp := &api.CalculatePrimeFactorsResponse{
				Result: n,
				Done:   false,
			}
			_ = stream.Send(resp)
		}
		_ = stream.Send(&api.CalculatePrimeFactorsResponse{
			Done: true,
		})
	}
	log.Println("Leaves CalculatePrimeFactors")

	return nil
}

// findPrimeFactors
func findPrimeFactors(c chan int64, num int64) {
	var i int64 = 2
	var numSqrt = int64(math.Sqrt(float64(num)))

	for ; i <= numSqrt; {
		if num%i == 0 {
			c <- i
			for num%i == 0 {
				num /= i
			}
		}

		i++
	}

	if num > 1 {
		c <- num
	}

	close(c)
}
