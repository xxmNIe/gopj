package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "server/proto"
	"server/srv"
)

const PORT = "8002"

func main() {
	lis, err := net.Listen("tcp",":"+PORT)
	if err !=nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s,&srv.Mserver{})
	pb.RegisterPeopleServer(s,&srv.Mserver{})
	if err := s.Serve(lis);err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
