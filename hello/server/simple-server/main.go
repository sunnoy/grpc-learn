/*
 *@Description
 *@author          lirui
 *@create          2021-05-08 14:34
 */
package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "openvpn/proto"
)

const port = ":1989"

type server struct {
	pb.UnimplementedGreeterServer
}

func (receiver server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "name:" + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)

	}

}
