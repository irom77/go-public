package main
//Based on https://gist.github.com/kotakanbe/d3059af990252ba89a82
import (
	"os/exec"
	"fmt"
	"time"
	"github.com/k0kubun/pp"
)

func ping(pingChan <-chan string, pongChan chan<- string) {
	for ip := range pingChan {
		_, err := exec.Command("ping", "-c 2", "-w 2", ip).Output()  //Linux
		//_, err := exec.Command("ping", "-n 2", "-w 2", ip).Output()
		if err == nil {
			pongChan <- ip
			//fmt.Printf("%s is alive\n", ip)
		} else {
			pongChan <- ""
			//fmt.Printf("%s is dead\n", ip)
		}
	}
}

func receivePong(pongNum int, pongChan <-chan string, doneChan chan<- []string) {
	var alives []string
	for i := 0; i < pongNum; i++ {
		ip := <-pongChan
		//fmt.Println("received: ", ip)
		alives = append(alives, ip)
	}
	doneChan <- alives
}

func list1s() []string { //Shield_Slice int
	res := make([]string, 256*64) //256*64
	for x := 192; x < 200; x++ {  //192-256
		for y := 0; y < 256; y++ {  //0-256
			res = append(res, fmt.Sprintf("10.%d.%d.1", x, y))
			//fmt.Printf("10.%d.%d.1", x, y)
		}
	}
	return res //[:Shield_Slice]
}

func main() {
	hosts := delete_empty(list1s())
	//fmt.Println(hosts, len(hosts))
	concurrentMax := 100
	pingChan := make(chan string, concurrentMax)
	pongChan := make(chan string, len(hosts))
	doneChan := make(chan []string)
	fmt.Printf("concurrentMax=%d hosts=%d -> %s...%s\n",concurrentMax, len(hosts),hosts[0], hosts[len(hosts)-1])
	start := time.Now()
	for i := 0; i < concurrentMax; i++ {
		go ping(pingChan, pongChan)
	}

	go receivePong(len(hosts), pongChan, doneChan)

	for _, ip := range hosts {
		pingChan <- ip
		//fmt.Println("sent: ", ip)
	}
	alives := <-doneChan
	result := delete_empty(alives)
	//fmt.Println(result)
	for _, ip := range result {
		fmt.Println(ip)
	}
	pp.Println(len(result))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func delete_empty (s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}