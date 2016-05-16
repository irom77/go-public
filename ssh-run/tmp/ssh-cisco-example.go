package tmp

import (
"flag"
"fmt"

//"github.com/ScriptRock/crypto/ssh"
"golang.org/x/crypto/ssh"
)

var (
USER = flag.String("user", "manager" , "ssh username")  //os.Getenv("USER")
HOST = flag.String("host", "127.0.0.1", "ssh server hostname")
PORT = flag.Int("port", 22, "ssh server port")
PASS = flag.String("pass", "", "ssh password")  //os.Getenv("SSH_PWD")
CMD =  flag.String("cmd", "", "command to run")
)

func init() { flag.Parse() }

func main() {
//ssh.Config -> Ciphers: []string{"aes128-cbc", "hmac-sha1", "none"},

sshConfig := &ssh.ClientConfig{
User: *USER,
Auth: []ssh.AuthMethod{ssh.Password(*PASS)},
//  HostKeyAlgorithms: []string{ssh.KeyAlgoRSA, ssh.KeyAlgoDSA},
}

addr := fmt.Sprintf("%s:%d", *HOST, *PORT)
client, err := ssh.Dial("tcp", addr, sshConfig)
if err != nil {
panic("Failed to dial: " + err.Error())
}

session, err := client.NewSession()
if err != nil {
client.Close()
panic("Failed to create session: " + err.Error())
}

out, err := session.CombinedOutput(*CMD)
if err != nil {
panic("Failed to run: " + err.Error())
}
fmt.Println(string(out))

// Close connection
client.Close()
}
