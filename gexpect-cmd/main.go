package main

import (
	"github.com/jamesharr/expect"
	"time"
	"fmt"
	"flag"
)

var (
	USER = flag.String("user", "manager", "ssh username") // or os.Getenv("USER") or os.Getenv("USERNAME")
	HOST = flag.String("host", "127.0.0.1", "ssh server name")
	PASS = flag.String("pass", "", "ssh password")
	CMD =  flag.String("cmd", "", "command to run")
)

func init() { flag.Parse() }

func main() {
	// Spawn an expect process
	ssh, err := expect.Spawn("ssh", "manager@10.29.1.65")
	ssh.SetTimeout(5 * time.Second)
	const PROMPT = `#` // `(?m)[^$]*$`

	// Login
	ssh.Expect(`[Pp]assword:`)
	ssh.SendMasked(*PASS) // SendMasked hides from logging
	ssh.Send("\n")
	ssh.Expect(PROMPT) // Wait for prompt

	// Run a command
	ssh.SendLn(*CMD)
	match, err := ssh.Expect(PROMPT) // Wait for prompt
	fmt.Println("command output:", match.Before)

	// Hit a timeout
	//ssh.SendLn("sleep 10") // This will cause a timeout
	match, err = ssh.Expect(PROMPT) // This will timeout
	if err == expect.ErrTimeout {
		fmt.Println("Session timed out.\n")
	}

	// Wait for EOF
	ssh.SendLn("logout")
	ssh.ExpectEOF()
}


