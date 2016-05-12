### Run ssh cmd on remote host 

```
C:\Users\irekromaniuk\Vagrant\trusty64\bin>ssh-cmd -h
Usage of ssh-cmd:
  -cmd string
        command to run
  -host string
        ssh server name (default "127.0.0.1")
  -pass string
        ssh password
  -user string
        ssh username (default "manager")
```


Radware Radware AppDirector throughput to Influx

See example, the throughput is last number of the command output

```
AppDirector#system last-sec-total-input

Total input on all ports in the last second (in Mbps): 137 

```

Tested on AppDirector Global v2.14.08DL