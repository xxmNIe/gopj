package server

import (
	"context"
	pb "hello-world/proto"
)

type helloService struct {
	pb.UnimplementedHelloWorldServer
}

func (h helloService) SayHelloWorld(ctx context.Context, request *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{
		Message: "test",
	},nil
}


func NewHelloService() *helloService{
	return &helloService{}

}





