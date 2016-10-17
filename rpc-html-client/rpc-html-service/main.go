package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/irom77/go-public/rpc-html-client/remote"
)

func main() {
	arith := new(remote.Arith)

	rpc.Register(arith)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	for {
		http.Serve(l, nil)
	}
}

