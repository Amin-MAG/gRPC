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

```makefile
$ make
```

But for this example we use go clients.

```makefile
$ protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      apis/prime.proto
```

## Server-Side codes

## Resource

[GitHub - ridha/grpc-streaming-demo: A quick demo of bi-directional streaming RPC's using grpc, go and python3](https://github.com/ridha/grpc-streaming-demo)