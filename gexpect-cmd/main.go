package main

import (
	"github.com/ThomasRooney/gexpect"
	"log"
	"flag"
	"fmt"
)

var (
	USER = flag.String("user", "manager", "ssh username") // or os.Getenv("USER") or os.Getenv("USERNAME")
	HOST = flag.String("host", "127.0.0.1", "ssh server name")
	PASS = flag.String("pass", "", "ssh password")
	CMD =  flag.String("cmd", "", "command to run")
)

func init() { flag.Parse() }

func main() {
log.Printf("Testing ssh... ")
fmt.Printf("%s %s %s %s", *USER, *PASS, *HOST, *CMD)
child, err := gexpect.Spawn("ssh " + *USER + "@" + *HOST)
if err != nil {
panic(err)
}

child.Expect("password:")
child.SendLine(*PASS)
child.Expect("#")
child.SendLine(*CMD)
child.Expect("#")
child.SendLine("logout")
log.Printf("\nSuccess\n")
}
