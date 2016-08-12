### gexpect to run commands on targets from file

##### EXAMPLE

goex-gaia -user="user@IP" -pass='password' -cmd="fw tab -t string_dictionary_table –x -y;fw fetch" -expert='password'

See example, the throughput is last number of the command output

```


```

Tested on 

######To bypass Host Key Checking
	$ cat /etc/ssh/ssh_config | grep StrictHostKeyChecking
	StrictHostKeyChecking no
	$ cat /etc/ssh/ssh_config | grep UserKnownHostsFile
	UserKnownHostsFile=/dev/null
	 