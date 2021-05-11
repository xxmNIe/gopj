package main

import (
	"fmt"
	pb "go-grpc-example/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type StreamService struct {
	pb.UnimplementedStreamServiceServer
}

func (s *StreamService) List(request *pb.StreamRequest, server pb.StreamService_ListServer) error {
	for n :=0;n<=6;n++ {
		err := server.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name: request.Pt.Name,
				Value: request.Pt.Value+int32(n),
			},
		})
		time.Sleep(500*time.Millisecond)
		if err !=nil{
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (s *StreamService) Record(server pb.StreamService_RecordServer) error {
	for {
		r,err := server.Recv()
		if err ==io.EOF {
			server.SendAndClose(&pb.StreamResponse{
				Pt: &pb.StreamPoint{
					Name: "gRPC Stream Server: Record",
					Value: 1,
				}})
		}
		if err !=nil {
			return err
		}
		log.Printf("stream.Recv pt.name: %s, pt.value: %d", r.Pt.Name, r.Pt.Value)
	}
	return nil
}

func (s *StreamService) Router(server pb.StreamService_RouterServer) error {
	panic("implement me")
}



const PORT1 = "9002"

func main() {

	//1.新建一个grpc  server
	server := grpc.NewServer()
	//注册
	pb.RegisterStreamServiceServer(server,&StreamService{})

	//tcp连接
	lis,err :=net.Listen("tcp",":"+PORT1)
	if err!=nil {
		fmt.Println("tcp listen err :",err)
	}
	//grpc挂到tcp连接上
	server.Serve(lis)


}
