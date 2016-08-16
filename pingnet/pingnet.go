package main
import (
	//"net"
	//"os/exec"
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

func pinger(targets []string, ch chan<- string)  {
	//defer wg.Done()
	for _, ip := range targets[0:] {
		/*_, err := exec.Command("ping", "-c", "1", "-w", "1", ip).Output()
		if err == nil {
		}*/
			ch <- fmt.Sprintf(ip)

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
	//targets := list1s()
	fmt.Printf("%s\n",len(list1s()))
	//ch := make(chan string)
	//go pinger(targets, ch)
	//go printer(ch)
	//wg.Wait()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	//os.Args[1]
}
