package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type args struct {
	A, B int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "10.29.21.208:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply int

	e := client.Call("Arith.Multiply", &args{4, 2}, &reply)
	if e != nil {
		log.Fatalf("Something went wrong: %s", err.Error())
	}

	fmt.Printf("The reply pointer value has been changed to: %d", reply)
}
