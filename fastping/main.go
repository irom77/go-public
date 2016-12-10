package main

import (
	"fmt"
	"time"
	"flag"
	"os"
	"bufio"
	"strconv"
	fastping "github.com/tatsushid/go-fastping"
	"net"
	"sync/atomic"
)

var (
	HOSTS = flag.String("a", "all", "destinations to ping, i.e. ./file.txt") // 'all', '/path/file' or i.e. '193'
	//CONCURRENTMAX = flag.Int("r", 200, "max concurrent pings")
	//PINGCOUNT = flag.String("c", "1", "ping count)")
	//PINGTIMEOUT = flag.Int("w", 1, "ping timout in s")
	version = flag.Bool("v", false, "Prints current version")
	PRINT = flag.Bool("print", true, "print to console")
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


func list1s(limit2 int) []string {
	//Shield_Slice int
	res := make([]string, 256 * 64) //256*64
	for x := 192; x < limit2; x++ {
		//192-256
		for y := 0; y < 256; y++ {
			//0-256
			res = append(res, fmt.Sprintf("10.%d.%d.1", x, y))
			//fmt.Printf("10.%d.%d.1", x, y)
		}
	}
	return res //[:Shield_Slice]
}

func main() {
	var hosts []string
	if *HOSTS == "all" {
		hosts = delete_empty(list1s(208))
		//fmt.Println(hosts, len(hosts))
	} else if num, err := strconv.Atoi(*HOSTS); err == nil {
		if (192 < num) && (num <= 208) {
			hosts = delete_empty(list1s(num))
		} else {
			hosts = delete_empty(list1s(208))
		}
	} else if pathExists(*HOSTS) {
		lines, err := readHosts(*HOSTS)
		hosts = delete_empty(lines)
		if err != nil {
			fmt.Println("Error reading file", *HOSTS)
		}
	} else {
		fmt.Println(*HOSTS)
		os.Exit(0)
	}

	fmt.Printf("hosts=%d -> %s...%s\n", len(hosts), hosts[0], hosts[len(hosts) - 1])
	start := time.Now()

	result := Ping(hosts)

	fmt.Printf("%.2fs %d/%d\n", time.Since(start).Seconds(),result,len(hosts))

}

// Ping takes a slice of IP addresses and return an int count of those that
// respond to ping.  The MaxRTT is set to 4 seconds.
func Ping(hosts []string) int {
	p := fastping.NewPinger()
	p.MaxRTT = 4 * time.Second
	var successCount, failCount uint64
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		atomic.AddUint64(&successCount, 1)
		fmt.Printf("IP Addr: %s receive, RTT: %v  successCount: %v \n", addr.String(), rtt, successCount)
	}
	p.OnIdle = func() {
		atomic.AddUint64(&failCount, 1)
		fmt.Println("timed out - finish")
	}

	for _, ip := range hosts {
		// fmt.Printf("adding ip: %v \n", ip)
		err := p.AddIP(ip)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error adding IP (%v): %v", ip, err)
		}
	}

	err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during Ping.Run(): %v", err)
	}

	return int(successCount)
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

func readHosts(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func pathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil { return true }
	if os.IsNotExist(err) { return false }
	return true
}