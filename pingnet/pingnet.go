package main
import (
	//"net"
	"os/exec"
	"fmt"
	//"os"
	"time"
	//"sync"
)

//var wg sync.WaitGroup

func list1s() []string {
	res := make([]string, 256*64)
	for x := 192; x < 193; x++ {  //192-156
		for y := 0; y < 256; y++ {
			res = append(res, fmt.Sprintf("10.%d.%d.1", x, y))
		}
	}
	return res
}

func pingip(ip string, ch chan<-string)  {
	//defer wg.Done()
	_, err := exec.Command("ping", "-c", "1", "-w", "1", ip).Output()
	if err == nil {
		ch <- fmt.Sprintf(ip)
	}
}

func main() {
	start := time.Now()
	targets := list1s()
	fmt.Printf("\n%d\n",len(targets))
	ch := make(chan string)
	for _,ip := range targets[0:]{
		//wg.Add(1)
		go pingip(ip, ch)
	}
	close(ch)
	for range targets[0:]{
		fmt.Println(<-ch)
	}
	//wg.Wait()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	//os.Args[1]
}
