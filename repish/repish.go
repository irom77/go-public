package main

import gexpect "github.com/ThomasRooney/gexpect"
import (
	"os"
	"fmt"
	"flag"
	"time"
	"net"
	"regexp"
)

var (
	USER = flag.String("u", "admin", "user name")
	HOST = flag.String("a", "localhost", "ip address")
	PROMPT1 = flag.String("p1", ">", "clish prompt")
	PROMPT2 = flag.String("p2", "]#", "expert prompt")
	PASS = flag.String("p", "", "clish password")
	EXPERT = flag.String("e", "", "expert password")
	CMD = flag.String("c", "fw stat", "command to run")
	INTERACT = flag.Bool("i", false, "interactive mode")
	TIMEOUT = flag.Int("t", 60, "timeout in sec")
	SEARCH = flag.String("s", "[>#]", "Search pattern in output")
	PORT = flag.String("webgui", "4434", "webgui port to test")
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

func main() {
	port := *PORT //1100/1400 webgui port
	match, status := RepishSocket(port)
	if  status != true {
		//log.Println("Can't connect")
		os.Exit(0)
	}
	if match == true {
		//log.Println("Connected")
		os.Exit(1)
	}
	searchPattern := *SEARCH  //i.e.`Done.` or 'WAN'
	var PROMPT string = *PROMPT1
	//log.Printf("ssh " + *USER + "@" + *HOST)
	//fmt.Println(*USER, *HOST, *PASS, *PROMPT1, *PROMPT2, *EXPERT, *CMD, *SEARCH, *INTERACT, *TIMEOUT)
	child, err := gexpect.Spawn("ssh " + *USER + "@" + *HOST)
	if err != nil {
		panic(err)
	}
	child.Expect("password:")
	child.SendLine(*PASS)
	child.Expect(PROMPT)
	if *EXPERT != "" {
		PROMPT = *PROMPT2
		child.SendLine("expert")
		child.Expect("password:")
		child.SendLine(*EXPERT)
		child.Expect(PROMPT)
	}
	child.SendLine(*CMD)
	if *INTERACT == true {
		child.Interact()
	} else {
		timeout := time.Duration(*TIMEOUT) * time.Second
		result, out, err := child.ExpectTimeoutRegexFindWithOutput(searchPattern, timeout)
		if err != nil {
			fmt.Printf("Error %v\nsearchPattern: %v\noutput: %v\nresult: %v\n", err, searchPattern, out, result)
		} else {
			fmt.Printf("searchPattern: %v\noutput: %v\nresult: %v\n", searchPattern, out, result)
		}
	}
	child.Close()
}

func RepishSocket(port string) (bool, bool) {
	match, _ := regexp.MatchString(":", *HOST)
	var socket string
	var status bool
	if match == true {
		socket = *HOST
	} else {
		socket = *HOST + ":" + port
	}
	dialer := &net.Dialer{Timeout: 2 * time.Second}
	conn, err := dialer.Dial("tcp", socket)
	if err != nil {
		//log.Println("Connection error:", err)
		status = false
	} else {
		//log.Println("Connected")
		defer conn.Close()
		status = true
	}
	return match, status
}