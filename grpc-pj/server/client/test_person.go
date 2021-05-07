package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "server/proto"
)


func main() {
	address     := "localhost:8002"
	conn,err := grpc.Dial(address,grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer  conn.Close()

	p := pb.NewPeopleClient(conn)
//
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	resp ,err :=p.SavePerson(context.Background(),&pb.SavePersonRequest{
		Person: &pb.Person{
			Name: "abc",
			Id: 1,
			Age: 2,
		},
	})
	if err != nil {
		log.Fatalf("could not save: %v", err)
	}
	log.Printf("saveperple: %s", resp.GetMessage())
	resp ,err =p.SavePerson(context.Background(),&pb.SavePersonRequest{
		Person: &pb.Person{
			Name: "abc",
			Id: 1,
			Age: 2,
		},
	})
	if err != nil {
		log.Fatalf("could not save: %v", err)
	}else {
		log.Printf("saveperple: %s", resp.GetMessage())
	}

}
