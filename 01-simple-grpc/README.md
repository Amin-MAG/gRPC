# Simple gRPC

Here we will implement a very basic server and clients. This server has a service responsible for calculating the majority of a list.

## Proto Files

```protobuf
syntax = "proto3";

option go_package = "bale_interview/apis";

package bale_interview;

service Majority {
  rpc CalculateListMajority(ListMajorityRequest) returns (ListMajorityReply) {}
}

// The request message containing the user's name.
message ListMajorityRequest {
  repeated int32 numbers = 1;
}

// The response message containing the greetings
message ListMajorityReply {
  string message = 1;
  int32 majorityNumber = 2;
}
```

`repeated` is for defining an array in proto files.

### Generating the Go files

You need to use `protoc` command to generate go files:

```bash
$ protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      apis/majority.proto
```

You specify the source, the output directory, and the proto file.  After executing this command some Go files will be generated.

## Server-side code

First you should create listener.

```go
lis, err := net.Listen("tcp", port)
if err != nil {
		log.Fatalf("failed to listen: %v", err)
}
```

Create the gRPC server. You can pass on some options if you want.

```go
s := grpc.NewServer()
```

Then you should register your services defined in proto files. First, You should pass on a type that implements that interface.

```go
type TestServer struct {
	*grpc.Server
	apis.UnimplementedMajorityServer
}

// CalculateListMajority to find most repeated number
// that repeated more than size / 2 times
func (s *TestServer) CalculateListMajority(ctx context.Context, in *apis.ListMajorityRequest) (*apis.ListMajorityReply, error) {
	integerCounter := make(map[int32]int32)
	listSize := int32(len(in.Numbers))

	// fill the integer counter
	fillIntegerCounter(in, integerCounter)

	// find the most repeated item
	exists, maxim := findMostRepeatedItemGreaterHalf(integerCounter, listSize)

	if exists == false {
		return &apis.ListMajorityReply{
			Message:        "Not found any majority",
			MajorityNumber: 0,
		}, nil
	} else {
		return &apis.ListMajorityReply{
			Message:        "majority is Founded",
			MajorityNumber: maxim,
		}, nil
	}
}
```

Here `TestServer` implements the interface. So we can use this struct to register.

```go
s := grpc.NewServer()
apis.RegisterMajorityServer(s, &TestServer{})
```

That's it.

## Client-Side code

It depends on what your client is. Here we assume we use Go language for clients too. This client uses the majority function to calculate the most repeated number.

```go
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
```