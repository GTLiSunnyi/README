package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"grpc-demo/proto"
)

type echoService struct {
	proto.UnimplementedEchoServiceServer
}

func (es *echoService) GetUnaryEcho(ctx context.Context, req *proto.EchoRequest) (*proto.EchoResponse, error) {
	res := "received: " + req.GetReq()
	fmt.Println(res)
	return &proto.EchoResponse{Res: res}, nil
}

func main() {
	rpc := grpc.NewServer()
	proto.RegisterEchoServiceServer(rpc, new(echoService))
	listener, err := net.Listen("tcp", "8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	rpc.Serve(listener)
}
