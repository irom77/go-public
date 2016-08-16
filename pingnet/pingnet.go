package main
//https://gist.github.com/kotakanbe/d3059af990252ba89a82
import (
	"net"
	"os/exec"
	//"github.com/k0kubun/pp"
	"fmt"
	"time"
	"os"
)

func ping(pingChan <-chan string, pongChan chan<- string) {
	for ip := range pingChan {
		_, err := exec.Command("ping", "-c 1", "-w 1", ip).Output()
		if err == nil {
			pongChan <- ip
		}
	}
}

func receivePong(pongNum int, pongChan <-chan string, doneChan chan<- []string) {
	var alives []string
	for i := 0; i < pongNum; i++ {
		ip := <-pongChan
		fmt.Println("received: ", ip)
		alives = append(alives, ip)
	}
	doneChan <- alives
}

func list1s(Shield_Slice int) []string {
	res := make([]string, 255) //256*64
	for x := 192; x < 256; x++ {  //192-256
		for y := 0; y < 256; y++ {
			res = append(res, fmt.Sprintf("10.%d.%d.1", x, y))
			//fmt.Printf("10.%d.%d.1", x, y)
		}
	}
	return res[:Shield_Slice]
}

func main() {
	hosts := delete_empty(list1s(300))
	fmt.Println(hosts)
	os.Exit(0)
	concurrentMax := 200
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
		//  fmt.Println("sent: " + ip)
	}
	alives := <-doneChan
	fmt.Println(alives)
	//pp.Println(len(alives))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
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