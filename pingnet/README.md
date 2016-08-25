## go to ping scan

##### EXAMPLES
cd $GOPATH/bin

$ go get -u github.com/irom77/go-public/pingnet

$ ./pingnet -h
Copyright 2016 @IrekRomaniuk. All rights reserved.
Usage of ./pingnet:
  -a string
        destinations to ping, i.e. ./file.txt (default "all")
  -c string
        ping count (-n 1 for Win) (default "-c 1")
  -print
        print to console (default true)
  -r int
        max concurrent pings (default 200)
  -v    Prints current version
  -w string
        ping timout in s (default "-w 1")
        
$ ./pingnet -a=200 -r=100 -print=false
concurrentMax=100 hosts=2048 -> 10.192.0.1...10.199.255.1
1094    
