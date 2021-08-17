# Simple gRPC bi-directional streaming
Here I want to create a simple server that calculates the prime factors of the given number. This connection is bi-directional streaming

## Defining the proto file

To specify that we want to use streaming we use `stream` keyword before the definition. In this example, I want to use bi-directional streaming. So I use the keyword before both the request and response.

```protobuf
syntax = "proto3";

option go_package = "prime/api";
package prime;

service Factors {
  // The stream keyword is specified before both the request type and response
  // type to make it as bidirectional streaming RPC method.
  rpc CalculatePrimeFactors (stream CalculatePrimeFactorsRequest) returns (stream CalculatePrimeFactorsResponse) {}
}

message CalculatePrimeFactorsRequest {
  int64 num = 1;
}

message CalculatePrimeFactorsResponse {
  int64 result = 1;
}
```

### Using a make file to generate the protofile

You can create a `Makefile`, when you want to generate the code in multiple languages.

```makefile
all: client server

client:
	@echo "--> Generating Kotlin client files"
	python3 -m grpc_tools.protoc -I protobuf/ --python_out=. --grpc_python_out=. protobuf/primefactor.proto
	@echo ""

server:
	@echo "--> Generating Go files"
	protoc -I protobuf/ --go_out=plugins=grpc:protobuf/ protobuf/primefactor.proto
	@echo ""

install:
	@echo "--> Installing Python grpcio tools"
	pip3 install -U grpcio grpcio-tools
	@echo ""
```

And then execute make command:

```bash
$ make
```

But for this example we use go clients.

```bash
$ protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      api/prime_factor.proto
```

## Server-Side code

First create a listener:

```go
lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
	log.Fatalln("ERROR:", err.Error())
}
```

Create the gRPC server. You can pass on some options if you want.

```go
s := grpc.NewServer()
```

For registering `CalculatePrimeFactors` we need to pass a struct that implements the interface.

```go
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
```

About stream functions:

- `stream.Recv()` will wait until some message comes from client.
- `io.EOF` will check if the stream has been closed or not.
- `stream.Send()` will send values to the client.
- When stream is closed, we should get out of this function.

## Client-Side code

Setup the connection:

```go
// Set up a connection to the server.
conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
if err != nil {
    log.Fatalln("cannot connect:", err)
}
defer func() {
    _ = conn.Close()
}()
```

Connect to the stream:

```go
// Setup the stream
c := api.NewFactorsClient(conn)
stream, err := c.CalculatePrimeFactors(context.Background())
if err != nil {
    log.Fatalln(err)
}
```

This stream is bi-directional so we should create a separate goroutine for receiving from server.

```go
{
	// In main function after creating stream
	wc := make(chan struct{})
	go receiveFromServer(wc, stream)
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
```

This function will receive each time something comes from server stream.

After this goroutine we implement an input mechanism to get number from user. Then send the numbers to the server using the `stream.Send()`.

```go
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
```

It will wait for the last receive and then iterate again. At the end we should close the stream:

```go
_ = stream.CloseSend()
time.Sleep(2 * time.Second)
```

> **Note: I don't know the reason but you don't use the `time.sleep()` the server is still waiting!**

## Resource

[GitHub - ridha/grpc-streaming-demo: A quick demo of bi-directional streaming RPC's using grpc, go and python3](https://github.com/ridha/grpc-streaming-demo)