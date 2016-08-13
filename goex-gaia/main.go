package main

import gexpect "github.com/ThomasRooney/gexpect"
import (
	"log"
	"os"
	"fmt"
	"flag"
	//"time"
	"time"
)

var (
	USERHOST = flag.String("user", "manager@localhost", "ssh user@host")
	PROMPT1 = flag.String("prompt1", ">", "clish prompt")
	PROMPT2 = flag.String("prompt2", "#", "expert prompt")
	PASS = flag.String("pass", "", "clish password")
	EXPERT = flag.String("expert", "", "expert password")
	CMD =  flag.String("cmd", "fw stat", "command to run")
	INTERACT = flag.Bool("interact", false, "interactive mode")
	TIMEOUT = flag.Int("timeout", 60, "timeout in sec")
	SEARCH = flag.String("searchPattern", "", "Search pattern in output")
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
	timeout := time.Duration(*TIMEOUT) * time.Second
	searchPattern := *SEARCH
	var PROMPT string = *PROMPT1
	log.Printf("ssh " + *USERHOST)
	//fmt.Println(*USERHOST, *PASS, *PROMPT1, *PROMPT2, *EXPERT, *CMD, *INTERACT)
	child, err := gexpect.Spawn("ssh " + *USERHOST)
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
	if *INTERACT {
		child.Interact()
	}
	result, out, err := child.ExpectTimeoutRegexFindWithOutput(searchPattern, timeout)
	if err != nil {
		fmt.Printf("Error %v\nsearchPattern: %v\noutput: %v\nresult: %v\n", err, searchPattern, out, result)
	}
	fmt.Printf("searchPattern: %v\noutput: %v\nresult: %v\n", searchPattern, out, result)

	//child.Expect(PROMPT)
	//child.Close()
}
