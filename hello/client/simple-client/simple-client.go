/*
 *@Description
 *@author          lirui
 *@create          2021-05-08 14:45
 */
package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "openvpn/proto"
)

const (
	address = "127.0.0.1:1989"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)

	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{
		Name: "hello world",
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)

	}

	log.Fatalf("greeting: %s", r.GetMessage())
}
