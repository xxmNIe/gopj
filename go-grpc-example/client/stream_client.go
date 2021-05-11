package main

import (
	"context"
	"fmt"
	pb "go-grpc-example/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

const PORT2 = "9002"

func main() {
	//监听
	conn,err := grpc.Dial(":"+PORT2,grpc.WithInsecure())

	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()
	//
	client := pb.NewStreamServiceClient(conn)
	//err = printLists(client,&pb.StreamRequest{Pt: &pb.StreamPoint{
	//	Name: "stream api client",Value: 2021}})
	//if err!=nil{
	//	fmt.Println("list err:",err)
	//}
	err = printRecord(client,&pb.StreamRequest{Pt: &pb.StreamPoint{
		Name: "stream api client",Value: 2021}})
	if err!=nil{
		fmt.Println("record err:",err)
	}
	err = printRoute(client,&pb.StreamRequest{Pt: &pb.StreamPoint{
		Name: "stream api printRoute",Value: 2021}})
	if err!=nil{
		fmt.Println("printRoute err:",err)
	}


}

func printLists(client pb.StreamServiceClient,r *pb.StreamRequest)error {
	stream,err := client.List(context.Background(),r)
	if err !=nil {
		return err
	}

	for {
		resp,err  :=stream.Recv()
		if err ==io.EOF {
			break
		}
		if err !=nil {
			return err
		}

		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}
	return nil
}
func printRecord(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}

	for n:=0;n<6;n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}

	resp,err :=stream.CloseAndRecv()
	if err != nil {
		return err
	}


	log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)

	return nil
}

func printRoute(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	return nil
}