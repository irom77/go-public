package main
import (
	//"net"
	"os/exec"
	"fmt"
	"os"
)

func list1s() []string {
	res := make([]string, 256*64)
	for x := 192; x < 256; x++ {
		for y := 0; y < 256; y++ {
			res = append(res, fmt.Sprintf("10.%d.%d.1", x, y))
		}
	}
	return res
}

func pingip(ip string) bool  {
	var alive bool
	_, err := exec.Command("ping", "-c", "1", "-w", "1", ip).Output()
	if err != nil {
		alive = false
	} else {
		alive = true
		fmt.Printf("Address %s is pingable", ip)
	}
	return alive
}

func main() {
	pingip(os.Args[1])
	//fmt.Printf("\n%v",list1s())
	fmt.Printf("\n%d\n",len(list1s()))
}
