package main
import (
	//"net"
	//"os/exec"
	"fmt"
	//"os"
	"time"
	//"sync"
	"os/exec"
)

//var wg sync.WaitGroup

func list1s() []string {
	res := make([]string, 255) //256*64
	for x := 192; x < 193; x++ {  //192-256
		for y := 0; y < 256; y++ {
			res = append(res, fmt.Sprintf("10.%d.%d.1", x, y))
			//fmt.Printf("10.%d.%d.1", x, y)
		}
	}
	return res
}

func pinger(targets []string, ch chan<- string)  {
	//defer wg.Done()
	for _, ip := range targets[0:] {
		_, err := exec.Command("ping", "-c", "1", "-w", "1", ip).Output()
		if err == nil {
			ch <- fmt.Sprintf(ip)
		}
	}
	//close(ch)
}

func printer (ch <-chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func main() {
	start := time.Now()
	targets := delete_empty(list1s())
	fmt.Printf("%d->%s...%s\n",len(targets),targets[0], targets[len(targets)-1])
	tmp(targets)
	//ch := make(chan string)
	//go pinger(targets, ch)
	//go printer(ch)
	//wg.Wait()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	//os.Args[1]
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

func tmp(targets []string) {
	for _, ip := range targets {
		fmt.Printf(ip)
	}
}
