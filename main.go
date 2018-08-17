package main

import (
	"context"
	"fmt"
	"github.com/utrack/grpc-connection-dropper/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	s := grpc.NewServer()
	l, err := net.Listen("tcp", "localhost:12349")
	if err != nil {
		panic(err)
	}
	fmt.Println(":12349")
	reflection.Register(s)

	sumpb.RegisterSummatorServer(s, &Handler{})

	err = s.Serve(l)
	if err != nil {
		panic(err)
	}

}

type Handler struct{}

func (s Handler) Sum(context.Context, *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Println("panicking")
	panic("test")
}
