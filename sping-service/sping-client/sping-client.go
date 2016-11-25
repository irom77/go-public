package main

import (
	"fmt"
	"log"
	"net/rpc"
	"flag"
	"os"
	"strconv"
	"bufio"
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
type args struct {
	hosts []string
	pcounter string
	ptimeout string
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
	fmt.Printf("hosts=%d -> %s...%s", len(hosts), hosts[0], hosts[len(hosts) - 1])
	fmt.Printf("\ntimeout=%ss counter=%s \n", *PINGTIMEOUT, *PINGCOUNT)
	//os.Exit(1)
	client, err := rpc.DialHTTP("tcp", "10.73.21.208:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply int

	e := client.Call("Ping.sping", &args{hosts, *PINGTIMEOUT, *PINGCOUNT}, &reply)
	if e != nil {
		log.Fatalf("Something went wrong: %s", err.Error())
	}

	fmt.Printf("The reply pointer value has been changed to: %d", reply)
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

