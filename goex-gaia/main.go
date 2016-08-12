package main

import (
	"github.com/jamesharr/expect"
	"time"
	"fmt"
	"flag"
)

var (
	USERHOST = flag.String("user", "manager@localhost", "ssh user@host")
	PROMPT1 = flag.String("clish prompt", ">", "prompt1")
	PROMPT2 = flag.String("expert prompt ", "#", "prompt2")
	PASS = flag.String("pass", "", "ssh password")
	EXPERT = flag.String("expert", "", "expert password")
	CMD =  flag.String("uptime", "", "command to run")
)

func init() { flag.Parse() }

func main() {
	//fmt.Println(*USERHOST, *PASS, *PROMPT, *EXPERT, *CMD)
	// Spawn an expect process
	ssh, _ := expect.Spawn("ssh", *USERHOST)
	ssh.SetTimeout(5 * time.Second)
	//const PROMPT = `#` // `(?m)[^$]*$`
	// Login
	ssh.Expect(`[Pp]assword:`)
	ssh.SendMasked(*PASS) // SendMasked hides from logging
	ssh.Send("\n")
	ssh.Expect(*PROMPT1) // Wait for prompt
	// Enter Expert mode
	if *EXPERT != "" {
		ssh.SendLn("expert\n")
		ssh.Expect(`[Pp]assword:`)
		ssh.SendMasked(*EXPERT)
		ssh.Send("\n")
		// Run a command
		ssh.Expect(*PROMPT2) // Wait for prompt
	}
	ssh.SendLn(*CMD)
	if *EXPERT != "" {
		match, _ := ssh.Expect(*PROMPT2) // Wait for prompt
		fmt.Println(match.Before)
	} else {
		match, _ := ssh.Expect(*PROMPT1)
		fmt.Println(match.Before)
	} // Wait for prompt}


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


