package main

import (
	"os/exec"
	"log"
)

type Target struct {
	Address string
	Replied bool
	ms string
}

type IP struct {
	Targets []Target

}

//List IP addresses of shield's IP address space
func (Addr *IP) Shields() {

}
//List n random IP addresses from shields's address space
func (Addr *IP) Randoms(n int ) {

}

//ping IP target address and keep result (Replied and ms)
func ping(Addr *IP) {
	for

	output, err := exec.Command("ping", "-c", "2", "8.8.8.8").Output()
	if err != nil {
		log.Fatal(err)
	}

}

func main () {

}