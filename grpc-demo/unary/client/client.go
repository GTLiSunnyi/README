package main

import (
	"bufio"
	"context"
	"grpc-demo/proto"
	"log"
	"os"

	"google.golang.org/grpc"
)

func main() {
	client, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	c := proto.NewEchoServiceClient(client)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		req := &proto.EchoRequest{Req: string(line)}
		res, err := c.GetUnaryEcho(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.GetRes())
	}
}
