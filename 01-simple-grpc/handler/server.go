package handler

import (
	"bale_interview/apis"
	"google.golang.org/grpc"
)

type TestServer struct {
	*grpc.Server
	apis.UnimplementedMajorityServer
}

// NewTestServer is to create new instance of TestServer
func NewTestServer() *TestServer {
	s := grpc.NewServer()

	apis.RegisterMajorityServer(s, &TestServer{})

	return &TestServer{
		Server: s,
	}
}
