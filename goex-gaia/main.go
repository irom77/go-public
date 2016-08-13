package main

import gexpect "github.com/ThomasRooney/gexpect"
import (
	"log"
	"os"
	"fmt"
	"flag"
)

var (
	USERHOST = flag.String("user", "manager@localhost", "ssh user@host")
	PROMPT1 = flag.String("prompt1", ">", "clish prompt")
	//PROMPT2 = flag.String("prompt2", "#", "expert prompt")
	PASS = flag.String("pass", "", "clish password")
	//EXPERT = flag.String("expert", "", "expert password")
	CMD =  flag.String("cmd", "fw stat", "command to run")
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

	log.Printf("ssh " + *USERHOST)

	child, err := gexpect.Spawn("ssh " + *USERHOST)
	if err != nil {
		panic(err)
	}
	child.Expect("password:")
	child.SendLine(*PASS)
	child.Expect(*PROMPT1)
	child.SendLine(*CMD)
	child.Interact()
	log.Printf("Success\n")
	child.Close()
}
