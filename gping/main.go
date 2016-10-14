package main
//go build -ldflags "-X main.Version=0.1 -X main.BuildTime=10/14/2016" github.com\irom77\go-public\gping
import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"bufio"
	"os/exec"
	"sync"
	"time"
	"runtime"
)

var (
	Version = "No Version Provided"
	BuildTime = ""
)

var (
	HOSTS = flag.String("a", "all", "destinations to ping, i.e. ./file.txt or '193'") // 'all', '/path/file' or i.e. '193'
	PINGCOUNT = flag.String("c", "1", "ping count")
	PINGTIMEOUT = flag.String("w", "1", "ping timout in s")
	version = flag.Bool("v", false, "Prints current version")
)

var (
count int
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
			fmt.Println("Error reading file: ", *HOSTS)
		}
	} else {
		fmt.Println("Input param error: ", *HOSTS)
		os.Exit(0)
	}
	var (
		os string
		timeout string
	)
	if runtime.GOOS == "windows" {
		fmt.Println("Windows OS detected")
		os = "-n"
		timeout = *PINGTIMEOUT + "000"
	}
	if runtime.GOOS == "linux" {    // also can be specified to FreeBSD
		fmt.Println("Unix/Linux type OS detected")
		os = "-c"
		timeout = *PINGTIMEOUT
	}
	fmt.Printf("hosts=%d -> %s...%s", len(hosts), hosts[0], hosts[len(hosts) - 1])
	fmt.Printf("\ntimeout=%sms %s counter=%s \n", *PINGTIMEOUT, os, *PINGCOUNT)
	//os.Exit(1)
	var wg sync.WaitGroup
	wg.Add(len(hosts))
	start := time.Now()
	runtime.GOMAXPROCS(MaxParallelism())
	for _, ip := range hosts {
		go ping(ip, &wg, os, timeout)
		//fmt.Println("sent: ", ip)
	}
	wg.Wait()
	//fmt.Printf("RESULT: %d/%d", count, len(hosts))
	fmt.Printf("RESULT: %d in %.2fs (%d CPUs)\n", count, time.Since(start).Seconds(),MaxParallelism())
}

func ping(ip string, wg *sync.WaitGroup, os string, timeout string ) {
	//_, err := exec.Command("ping", "-c 1", "-w 1", ip).Output()  //Linux
	//result , err := exec.Command("ping", *PINGCOUNT, *PINGTIMEOUT, ip).Output()
	_ , err := exec.Command("ping", os, *PINGCOUNT, "-w", timeout, ip).Output()
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

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}
