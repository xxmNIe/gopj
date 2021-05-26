package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(2)
	for i:=0;i<10;i++ {
		fmt.Println(r.Next())
	}

}
