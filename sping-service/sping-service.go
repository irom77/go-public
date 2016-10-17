package main
//go build -ldflags "-X main.Version=0.1 -X main.BuildTime=10/14/2016" github.com\irom77\go-public\gping
//discussed http://stackoverflow.com/questions/40049884/golang-sync-waitgroup-doesnt-complete-on-linux/40051153#40051153
import (
	"fmt"
	"os"
	"os/exec"
	"sync"
	"time"
	"runtime"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

var count int
type Args struct {
	hosts []string
	pcounter string
	ptimeout string
}
type Ping int

func (t *Ping) sping( args *Args, reply *int ) error {
	fmt.Printf("Args received: %+v\n", args)
	var (
		os string
		timeout string
	)
	if runtime.GOOS == "windows" {
		fmt.Println("Windows OS detected")
		os = "-n"
		timeout = Args.ptimeout + "000"
	}
	if runtime.GOOS == "linux" {    // also can be specified to FreeBSD
		fmt.Println("Unix/Linux type OS detected")
		os = "-c"
		timeout = Args.ptimeout
	}
	var wg sync.WaitGroup
	wg.Add(len(Args.hosts))
	start := time.Now()
	runtime.GOMAXPROCS(MaxParallelism())
	for _, ip := range Args.hosts {
		go ping(ip, &wg, os, timeout, Args.pcounter)
		//fmt.Println("sent: ", ip)
	}
	wg.Wait()
	//fmt.Printf("RESULT: %d/%d", count, len(hosts))
	fmt.Printf("RESULT: %d in %.2fs (%d CPUs)\n", count, time.Since(start).Seconds(),MaxParallelism())
	*reply = count
	return nil
}

func main() {
	ping := new(Ping)

	rpc.Register(ping)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	for {
		http.Serve(l, nil)
	}
}

func ping(ip string, wg *sync.WaitGroup, os string, timeout string, pcounter string ) {
	//_, err := exec.Command("ping", "-c 1", "-w 1", ip).Output()  //Linux
	//result , err := exec.Command("ping", *PINGCOUNT, *PINGTIMEOUT, ip).Output()
	_ , err := exec.Command("ping", os, pcounter, "-w", timeout, ip).Output()
	//_, err := exec.Command("ping", "-n 1", "-w 1", ip).Output()
	//fmt.Printf("%s\n", result)
	if err == nil {
		count++
		fmt.Printf("%d %s \n", count, ip)
	} else {
		//fmt.Printf("%s is dead\n", ip)
	}
	wg.Done()
}


func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}


