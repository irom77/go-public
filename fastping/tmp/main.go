package main

/*
irek@nms01m bin]$ ./fastping.sh pingsample.txt
Wed Aug 17 15:26:06 EDT 2016
IP Addr: 10.197.57.1 receive, RTT: 26.111138ms
IP Addr: 10.197.69.1 receive, RTT: 43.331557ms
IP Addr: 10.192.66.1 receive, RTT: 99.566382ms
IP Addr: 10.197.71.1 receive, RTT: 48.479079ms
IP Addr: 8.8.8.8 receive, RTT: 9.400096ms
IP Addr: 10.197.90.1 receive, RTT: 23.394716ms
*/


import (
	"fmt"
	"net"
	"time"
	"github.com/tatsushid/go-fastping"
	"os"
)

func main() {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		//fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func list1s(limit2 int) []string {
	res := make([]string, 256 * 64) //256*64
	for x := 192; x < limit2; x++ {
		//192-256
		for y := 0; y < 256; y++ {
			//0-256
			res = append(res, fmt.Sprintf("10.%d.%d.1", x, y))
			//fmt.Printf("10.%d.%d.1", x, y)
		}
	}
	return res
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}


