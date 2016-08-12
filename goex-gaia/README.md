### go expect to run commands on Checkpoint Gaia

##### EXAMPLE
cd $GOPATH\bin
go get -u github.com/irom77/go-public/goex-gaia
go build -ldflags "-X main.BuildTime=`date -u +.%Y%m%d.%H%M%S` -X main.Version=1.0.1" github.com/irom77/go-public/goex-gaia
goex-gaia -user="user@IP" -pass='password' -cmd="df -kh" -expert='password'
goex-gaia -user="user@IP" -pass='password' -cmd="show software-version" 

######To bypass Host Key Checking
	$ cat /etc/ssh/ssh_config | grep StrictHostKeyChecking
	StrictHostKeyChecking no
	$ cat /etc/ssh/ssh_config | grep UserKnownHostsFile
	UserKnownHostsFile=/dev/null
	 