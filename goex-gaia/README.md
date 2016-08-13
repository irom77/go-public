### go expect to run commands on Checkpoint Gaia

##### EXAMPLE
cd $GOPATH/bin
go get -u github.com/irom77/go-public/goex-gaia
go build -ldflags "-X main.BuildTime=`date -u +.%Y%m%d.%H%M%S` -X main.Version=1.0.1" github.com/irom77/go-public/goex-gaia
goex-gaia -user="user@IP" -pass='password' -cmd="df -kh" -expert='password'
goex-gaia -user="user@IP" -pass='password' -cmd="ip addr" -expert='password'
./goex-gaia -user='admin@10.199.107.1' -pass='password' -prompt1='#' -cmd='fw tab -t string_dictionary_table –x'
goex-gaia -user="user@IP" -pass='password' -cmd="show software-version" 
goex-gaia -user="user@IP" -pass='password' <- default cmd is 'fw stat'
$ ./goex-gaia -user="admin@10.198.2.1" -pass='password' -expert='password' -cmd="fw fetch" -search='Done.'


To read 'userhost' from file: 
```
$ cat goex-gaia.sh
#!/bin/bash
date
cat pinglist.txt |  while read output
do
    USERHOST="$1@$output"
    echo $USERHOST
    ./goex-gaia -user=$USERHOST -pass='password'
done
$./goex-gaia.sh admin
```


######To bypass Host Key Checking
	$ cat /etc/ssh/ssh_config | grep StrictHostKeyChecking
	StrictHostKeyChecking no
	$ cat /etc/ssh/ssh_config | grep UserKnownHostsFile
	UserKnownHostsFile=/dev/null
	 