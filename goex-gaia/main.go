package main

import gexpect "github.com/ThomasRooney/gexpect"
import "log"

func main() {
	log.Printf("Testing ssh... ")

	child, err := gexpect.Spawn("ssh admin@10.199.107.1")
	if err != nil {
		panic(err)
	}
	child.Expect("password:")
	child.SendLine("n3w@y!n")
	child.Expect("#")
	child.SendLine("fw fetch")
	//child.Interact()
	child.SendLine("exit")
	log.Printf("Success\n")
	child.Close()
}
