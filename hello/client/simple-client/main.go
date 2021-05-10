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
	"time"
)

const (
	address     = "127.0.0.1:1989"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)

	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	name := defaultName

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{
		Name: name,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)

	}

	log.Fatalf("greeting: %s", r.GetMessage())
}
