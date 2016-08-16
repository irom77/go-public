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
	USER = flag.String("user", "admin", "user name")
	HOST = flag.String("addr", "localhost", "ip address")
	PROMPT1 = flag.String("prompt1", ">", "clish prompt")
	PROMPT2 = flag.String("prompt2", "#", "expert prompt")
	PASS = flag.String("pass", "", "clish password")
	EXPERT = flag.String("expert", "", "expert password")
	CMD = flag.String("cmd", "fw stat", "command to run")
	INTERACT = flag.Bool("interact", false, "interactive mode")
	TIMEOUT = flag.Int("timeout", 60, "timeout in sec")
	SEARCH = flag.String("search", "", "Search pattern in output")
	PORT = flag.String("port", "4434", "webgui port to test")
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
	colon, status := RepishSocket(port)
	if  status != true {
		//log.Println("Can't connect")
		os.Exit(0)
	}
	if colon == true {
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
	child.Close()
	Search(*CMD, searchPattern)
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

func Search(cli, searchPattern string) {
	p, err := gexpect.Spawn(cli)
	if err != nil {
		panic(err)
	}
	if *INTERACT == true {
		p.Interact()
	}
	if searchPattern != "" {
		timeout := time.Duration(*TIMEOUT) * time.Second
		result, out, err := p.ExpectTimeoutRegexFindWithOutput(searchPattern, timeout)
		if err != nil {
			fmt.Printf("Error %v\nsearchPattern: %v\noutput: %v\nresult: %v\n", err, searchPattern, out, result)
		} else {
			fmt.Printf("searchPattern: %v\noutput: %v\nresult: %v\n", searchPattern, out, result)
		}
	}
	p.Close()
}