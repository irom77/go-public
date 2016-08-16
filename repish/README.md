## go expect to run commands on Checkpoint Gaia

##### EXAMPLES
cd $GOPATH/bin

go get -u github.com/irom77/go-public/repish

go build -ldflags "-X main.BuildTime=`date -u +.%Y%m%d.%H%M%S` -X main.Version=1.0.1" github.com/irom77/go-public/goex-gaia

goex-gaia -user="user@IP" -pass='' -cmd="df -kh" -expert=''

goex-gaia -user="user@IP" -pass='' -cmd="ip addr" -expert=''

goex-gaia -user='admin@10.199.107.1' -pass='' -prompt1='#' -cmd='fw tab -t string_dictionary_table –x -y'

goex-gaia -user="user@IP" -pass='' -cmd="show software-version" 

goex-gaia -user="user@IP" -pass='' <- default cmd is 'fw stat'

./repish -addr="10.197.57.1" -user="admin" -pass='pass' -cmd="add user indeni type admin password pass permission RW" -search='>'

```
searchPattern: >
output:  add user indeni type admin password pass permission RW
ADVHMASAFEPORT197-57>
```

goex-gaia -user="admin@10.198.2.1" -pass='' -expert='' -cmd="fw fetch" -search='Done.'

goex-gaia -user="admin@10.199.16.1" -pass='' -cmd="show software-version" -search=' - '

goex-gaia -user="admin@10.199.16.1" -pass='' -expert='' -cmd="fw tab -t string_dictionary_table -x -y" -search='Clearing'

$ goex-gaia -user="admin@10.199.16.1" -pass='' -expert='' -cmd="fw tab -t string_dictionary_table -x -y; fw fetch" -search='Done.'

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
./goex-gaia -user="admin@10.199.107.1" -pass='' -cmd="fw tab -t string_dictionary_table -x -y; fw fetch" -search='Done.' -prompt1='#' -timeout='120'

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
$ cat goex-gaia.sh
#!/bin/bash
date
cat pinglist.txt |  while read output
do
    USERHOST="$1@$output"
    echo $USERHOST
    ./goex-gaia -user=$USERHOST -pass=''
done
$./goex-gaia.sh admin
```


######To bypass Host Key Checking

	$ cat /etc/ssh/ssh_config | grep StrictHostKeyChecking
	StrictHostKeyChecking no
	$ cat /etc/ssh/ssh_config | grep UserKnownHostsFile
	UserKnownHostsFile=/dev/null
	 