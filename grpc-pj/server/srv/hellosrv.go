package srv

import (
	"context"
	"errors"
	"server/db"
	pb "server/proto"
)

type Mserver struct {
	pb.UnimplementedPeopleServer
	pb.UnimplementedGreeterServer

}

func (m *Mserver) SavePerson(ctx context.Context, request *pb.SavePersonRequest) (*pb.SavePersonReplay, error) {
	p := request.GetPerson()

	if _,ok :=db.Get(int(p.GetId()));ok {
		return &pb.SavePersonReplay{Message: ""},errors.New("id已存在")
	}
	db.Add(&db.Person{
		Id: int(p.GetId()),
		Name: p.GetName(),
		Age: int(p.GetAge()),
		})
	return &pb.SavePersonReplay{Message: "添加成功"},nil
}

func (m *Mserver) SavePersons(ctx context.Context, request *pb.SavePersonListRequest) (*pb.SavePersonListReplay, error) {
	ps := request.GetPeoples()
	var count int64 =0
	for _,p :=range ps {
		if db.Add(&db.Person{
			Id: int(p.GetId()),
			Name: p.GetName(),
			Age: int(p.GetAge()),
		}) {
			count++
		}
	}
	return &pb.SavePersonListReplay{Count: count},nil
}



func (m *Mserver) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReplay, error) {
	return &pb.HelloReplay{
		Message: "hello"+request.GetName(),
	},nil
}


