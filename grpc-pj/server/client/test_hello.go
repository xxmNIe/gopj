package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	pb "server/proto"
	"time"
)

const (
	address     = "localhost:8002"
	defaultName = "world"
)
func main() {
	conn,err := grpc.Dial(address,grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer  conn.Close()

	c := pb.NewGreeterClient(conn)

	name := defaultName

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx,cancle := context.WithTimeout(context.Background(),time.Second)
	defer cancle()

	r ,err := c.SayHello(ctx,&pb.HelloRequest{Name:name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
