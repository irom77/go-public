package main

import (
	"golang.org/x/crypto/ssh"
	"fmt"
	"bytes"
	//"os"
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
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig.
	config := &ssh.ClientConfig{
		User: *USER,
		Auth: []ssh.AuthMethod{
			ssh.Password(*PASS),
		},
		/*Config: ssh.Config{
			//Ciphers: []string{"3des-cbc", "blowfish-cbc", "arcfour"},
			//Ciphers: ssh.AllSupportedCiphers(), // include cbc ciphers
			//or edit GOPATH/src/golang.org/x/crypto/ssh/common.go

		},*/
	}
	client, err := ssh.Dial("tcp", *HOST + ":22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(*CMD); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}