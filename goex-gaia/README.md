### gexpect to run commands on targets from file

##### EXAMPLE

gex-run -user="user@IP" -pass='password' -cmd="show software-version" -prompt=">"

See example, the throughput is last number of the command output

```


```

Tested on 

######To bypass Host Key Checking
	$ cat /etc/ssh/ssh_config | grep StrictHostKeyChecking
	StrictHostKeyChecking no
	$ cat /etc/ssh/ssh_config | grep UserKnownHostsFile
	UserKnownHostsFile=/dev/null
	 