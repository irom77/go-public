package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"fmt"
)

// Args is a data structure for the incoming arguments
// This needs to be exported for the RPC to be valid/work
type Args struct {
	A, B int
}

// Arith is our functions return type
// This also needs to be exported
type Arith int

// Multiply does simply multiplication on provided arguments
// This also needs to be exported
func (t *Arith) Multiply(args *Args, reply *int) error {
	fmt.Printf("Args received: %+v\n", args)
	*reply = args.A * args.B
	return nil
}
func rpcservice() {
	arith := new(Arith)

	rpc.Register(arith)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1235")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	for {
		http.Serve(l, nil)
	}
}


