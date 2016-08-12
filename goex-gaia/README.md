### go expect to run commands on Checkpoint Gaia

##### EXAMPLE

goex-gaia -user="user@IP" -pass='password' -cmd="df -kh" -expert='password'
goex-gaia -user="user@IP" -pass='password' -cmd="show software-version" 

######To bypass Host Key Checking
	$ cat /etc/ssh/ssh_config | grep StrictHostKeyChecking
	StrictHostKeyChecking no
	$ cat /etc/ssh/ssh_config | grep UserKnownHostsFile
	UserKnownHostsFile=/dev/null
	 