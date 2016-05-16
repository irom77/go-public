### gexpect to run single command

##### EXAMPLE

Radware Radware AppDirector throughput to Influx

gex-run -user="user@IP" -pass='password' -cmd="system last-sec-total-input" -prompt="#" | grep -E '[0-9]{1}' | awk '{print $NF}'

See example, the throughput is last number of the command output

```
AppDirector#system last-sec-total-input

Total input on all ports in the last second (in Mbps): 137 

```

Tested on AppDirector Global v2.14.08DL