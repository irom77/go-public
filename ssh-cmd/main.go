package main

import (
	"golang.org/x/crypto/ssh"
	"fmt"
	"bytes"
	"flag"
	//"log"
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
	// A public key may be used to authenticate against the remote
	// server by using an unencrypted PEM-encoded private key file.
	//
	// If you have an encrypted private key, the crypto/x509 package
	// can be used to decrypt it.
	/*key, err := ioutil.ReadFile("/home/irek/.ssh/id_rsa")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}*/
	config := &ssh.ClientConfig{
		User: *USER,
		Auth: []ssh.AuthMethod{
			ssh.Password(*PASS),
			// Use the PublicKeys method for remote authentication.
			//ssh.PublicKeys(signer),
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
	// Set up terminal modes
	/*modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}*/
	// Request pseudo terminal
	/*if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		log.Fatalf("request for pseudo terminal failed: %s", err)
	}*/
	// Start remote shell
	/*if err := session.Shell(); err != nil {
		log.Fatalf("failed to start shell: %s", err)
	}*/
	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(*CMD); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}