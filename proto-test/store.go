package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	pb "ptoto-test/proto"
)

func main() {
	store()
	load()
}

var fname string = "books"
func store(){

	book := &pb.AddressBook{People: []*pb.Person{&pb.Person{
		Name: "gty",
		Id: 250,
		Email: "250@qq.com",
	}}}
	out,err := proto.Marshal(book)

	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

func load(){
	in,err :=ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	book :=&pb.AddressBook{}
	if err := proto.Unmarshal(in,book);err !=nil {
		log.Fatalln("Error decode file:", err)
	}
	fmt.Println(book)
}