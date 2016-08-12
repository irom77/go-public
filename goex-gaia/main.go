package main

import (
	"github.com/jamesharr/expect"
	"time"
	"fmt"
	"flag"
)

var (
	USERHOST = flag.String("user", "manager@localhost", "ssh user@host")
	PROMPT = flag.String("prompt", "#", "ssh prompt")
	PASS = flag.String("pass", "", "ssh password")
	EXPERT = flag.String("expert", "", "expert password")
	CMD =  flag.String("cmd", "", "command to run")
)

func init() { flag.Parse() }

func main() {
	// Spawn an expect process
	ssh, _ := expect.Spawn("ssh", *USERHOST)
	ssh.SetTimeout(5 * time.Second)
	//const PROMPT = `#` // `(?m)[^$]*$`
	// Login
	ssh.Expect(`[Pp]assword:`)
	ssh.SendMasked(*PASS) // SendMasked hides from logging
	ssh.Send("\n")
	ssh.Expect(*PROMPT) // Wait for prompt
	// Enter Expert mode
	ssh.SendLn("expert\n")
	ssh.Expect(`[Pp]assword:`)
	ssh.SendMasked(*EXPERT)
	// Run a command
	ssh.Expect(*PROMPT) // Wait for prompt
	ssh.SendLn(*CMD)
	match, _ := ssh.Expect("#") // Wait for prompt
	fmt.Println(match.Before)

	// Hit a timeout
	//ssh.SendLn("sleep 10") // This will cause a timeout
	//ssh.Expect(PROMPT) // This will timeout
	/*if err == expect.ErrTimeout {
		fmt.Println("Session timed out.\n")
	}*/
	ssh.Close();
	// Wait for EOF
	/*ssh.SendLn("logout")
	ssh.ExpectEOF()*/
}


