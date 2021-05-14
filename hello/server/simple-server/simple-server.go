/*
 *@Description
 *@author          lirui
 *@create          2021-05-08 14:34
 */
package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	pb "openvpn/proto"
)

const port = ":1989"

type server struct {
	pb.UnimplementedGreeterServer
}

func (receiver server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "name:" + in.GetName() + "server reply"}, nil
}

func main() {

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})

	reflection.Register(s)

	lis, _ := net.Listen("tcp", port)

	s.Serve(lis)

}
