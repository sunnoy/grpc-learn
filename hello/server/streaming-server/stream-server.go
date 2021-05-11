/*
 *@Description
 *@author          lirui
 *@create          2021-05-08 17:46
 */
package main

import (
	"bufio"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	pb "openvpn/proto"
	"os/exec"
)

type StreamService struct {
	pb.UnimplementedStreamServiceServer
}

const PORT = ":1989"

func main() {
	server := grpc.NewServer()

	// 第二个参数类型是 StreamServiceServer interface 而 StreamService 实现了 StreamServiceServer 接口
	pb.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("list %s", err)

	}

	// Serve accepts incoming connections on the listener lis, creating a new
	// ServerTransport and service goroutine for each. The service goroutines
	// read gRPC requests and then call the registered handlers to reply to them.
	server.Serve(lis)

}

func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {

	////////////////////////////////////////////////////
	//for n := 0; n <= 10; n++ {
	//
	//	time.Sleep(time.Second * 1)
	//
	//	// 开始发送消息
	//	err := stream.Send(&pb.StreamResponse{
	//		Pt: &pb.StreamPoint{
	//			Name:  r.Pt.Name,
	//			Value: r.Pt.Value + int32(n),
	//		},
	//	})
	//
	//	if err != nil {
	//		return err
	//	}
	//}

	/////////////////////////////////////////////////////
	cmd := exec.Command("ping", "127.0.0.1")

	stdout, _ := cmd.StdoutPipe()

	cmd.Start()

	buf := bufio.NewReader(stdout)

	for {
		line, _, _ := buf.ReadLine()

		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  string(line),
				Value: int32(100),
			},
		})

		if err != nil {
			return err
		}

	}

	return nil
}

func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{Pt: &pb.StreamPoint{Name: "gRPC Stream Server: Record", Value: 1}})

		}
		if err != nil {
			return err
		}
		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)

	}

	return nil
}

func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	n := 0
	for {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  "gPRC Stream Client: Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}

		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++

		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}

	return nil
}
