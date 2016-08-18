## go expect to run commands on Checkpoint Gaia

##### EXAMPLES
cd $GOPATH/bin

go get -u github.com/irom77/go-public/repish

go build -ldflags "-X main.BuildTime=`date -u +.%Y%m%d.%H%M%S` -X main.Version=1.0.1" github.com/irom77/go-public/repsih

repish -a="IP" -p='' -c="df -kh" -e=''

repsih -a="IP" -p='' -c="ip addr" -e=''

repsih -a="10.199.107.1" -p='' -p1='#' -c='fw tab -t string_dictionary_table –x -y'

repsih -a="IP" -p='' -c="show software-version" 

repsih -a="IP" -p='' <- default cmd is 'fw stat'

./repish -a="10.197.57.1" -u="admin" -p='pass' -c="add user indeni type admin password pass permission RW" -s='>'

```
searchPattern: >
output:  add user indeni type admin password pass permission RW
ADVHMASAFEPORT197-57>
```

repsih -u="10.198.2.1" -p='' -e='' -c="fw fetch" -s='Done.'

repsih -u="10.199.16.1" -p='' -c="show software-version" -s=' - '

repsih -u="10.199.16.1" -p='' -e='' -c="fw tab -t string_dictionary_table -x -y" -s='Clearing'

$ repsih -a="10.199.16.1" -p='' -e='' -c="fw tab -t string_dictionary_table -x -y; fw fetch" -s='Done.'

```
2016/08/14 08:53:58 ssh admin@10.199.16.1
searchPattern: Done.
output:  
fw tab -t string_dictionary_table -x -y; fw fetch
[Expert@199-16]# fw tab -t string_dictionary_table -x -y; fw fetch
[Expert@199-16]# fw tab -t string_dictionary_table -x -y; fw fetch 
Clearing table string_dictionary_table
Fetching Security Policy From: 1.1.1.1

Local Security Policy is Up-To-Date.

Installing Security Policy...

Done.
result: [Done.]
```

and for bash user (no expert mode):
./repsih -a="10.199.107.1" -p='' -c="fw tab -t string_dictionary_table -x -y; fw fetch" -search='Done.' -p1='#' -t='120'

```
2016/08/14 09:16:28 ssh admin@10.199.107.1
searchPattern: Done.
output:  fw tab -t string_dictionary_table -x -y; fw fetch

Clearing table string_dictionary_table
Fetching Security Policy from '216.57.142.237'

Local Security Policy is Up-To-Date.

Installing Security Policy...

Done.
result: [Done.]
```

$  ./repish -a="10.197.57.1" -p='' -e='' -s='Interrupt' -c="ifconfig WAN" | grep addr: | awk '{ print $2 }'

```
addr:192.168.1.150
```


To read 'userhost' from file: 

```
$ cat repsih.sh
#!/bin/bash
date
cat pinglist.txt |  while read output
do
    USERHOST="$1@$output"
    echo $USERHOST
    ./repsih -user=$USERHOST -pass=''
done
$./repsih.sh admin
```


######To bypass Host Key Checking

	$ cat /etc/ssh/ssh_config | grep StrictHostKeyChecking
	StrictHostKeyChecking no
	$ cat /etc/ssh/ssh_config | grep UserKnownHostsFile
	UserKnownHostsFile=/dev/null
	 