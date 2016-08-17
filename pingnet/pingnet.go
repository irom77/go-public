package main
//Based on https://gist.github.com/kotakanbe/d3059af990252ba89a82
import (
	"os/exec"
	"fmt"
	"time"
	"github.com/k0kubun/pp"
	"flag"
	"os"
)
var (
	HOSTS = flag.String("a", "all", "destinations to ping") // 'all', '/path/file' or i.e. '193'
	CONCURRENTMAX = flag.Int("r", 200, "max concurrent pings")
	PINGOPTIONS = flag.String("o", "-c 1 -w 1", "ping options")
	version = flag.Bool("v", false, "Prints current version")
)
var (
	Version = "No Version Provided"
	BuildTime = ""
)
func init() {
	flag.Usage = func() {
		fmt.Printf("Copyright 2016 @IrekRomaniuk. All rights reserved.\n")
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *version {
		fmt.Printf("App Version: %s\nBuild Time : %s\n", Version, BuildTime)
		os.Exit(0)
	}
}
func ping(pingChan <-chan string, pongChan chan<- string) {
	for ip := range pingChan {
		_, err := exec.Command("ping", "-c 1", "-w 1", ip).Output()  //Linux
		//_, err := exec.Command("ping", "-n 1", "-w 1", ip).Output()
		if err == nil {
			pongChan <- ip
			//fmt.Printf("%s is alive\n", ip)
		} else {
			pongChan <- ""
			//fmt.Printf("%s is dead\n", ip)
		}
	}
}

func receivePong(pongNum int, pongChan <-chan string, doneChan chan<- []string) {
	var alives []string
	for i := 0; i < pongNum; i++ {
		ip := <-pongChan
		//fmt.Println("received: ", ip)
		alives = append(alives, ip)
	}
	doneChan <- alives
}

func list1s() []string { //Shield_Slice int
	res := make([]string, 256*64) //256*64
	for x := 192; x < 200; x++ {  //192-256
		for y := 0; y < 256; y++ {  //0-256
			res = append(res, fmt.Sprintf("10.%d.%d.1", x, y))
			//fmt.Printf("10.%d.%d.1", x, y)
		}
	}
	return res //[:Shield_Slice]
}

func main() {
	var hosts []string
	if *HOSTS == "all" {
		hosts = delete_empty(list1s())
		//fmt.Println(hosts, len(hosts))
	} else if *HOSTS == "test"{
		fmt.Println("Test")
	}

	concurrentMax := *CONCURRENTMAX
	pingChan := make(chan string, concurrentMax)
	pongChan := make(chan string, len(hosts))
	doneChan := make(chan []string)
	fmt.Printf("concurrentMax=%d hosts=%d -> %s...%s\n",concurrentMax, len(hosts),hosts[0], hosts[len(hosts)-1])
	start := time.Now()
	for i := 0; i < concurrentMax; i++ {
		go ping(pingChan, pongChan)
	}

	go receivePong(len(hosts), pongChan, doneChan)

	for _, ip := range hosts {
		pingChan <- ip
		//fmt.Println("sent: ", ip)
	}
	alives := <-doneChan
	result := delete_empty(alives)
	//fmt.Println(result)
	for _, ip := range result {
		fmt.Println(ip)
	}
	pp.Println(len(result))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func delete_empty (s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}