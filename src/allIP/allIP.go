package allIP

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
)

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

type Pong struct {
	Ip    string
	Alive bool
}

func ping(pingChan <-chan string, pongChan chan<- Pong) {
	for ip := range pingChan {
		_, err := exec.Command("ping", "-c1", "-t1", ip).Output()
		var alive bool
		if err != nil {
			alive = false
		} else {
			alive = true
		}
		pongChan <- Pong{Ip: ip, Alive: alive}
	}
}

func receivePong(pongNum int, pongChan <-chan Pong, doneChan chan<- []Pong) {
	var alives []Pong
	for i := 0; i < pongNum; i++ {
		pong := <-pongChan
		//  fmt.Println("received:", pong)
		if pong.Alive {
			alives = append(alives, pong)
		}
	}
	doneChan <- alives
}

func AllIP() {
	var h string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
				h = ipnet.IP.String()
			}
		}
	}
	hosts, _ := Hosts(h + "/21")
	concurrentMax := 100
	pingChan := make(chan string, concurrentMax)
	pongChan := make(chan Pong, len(hosts))
	doneChan := make(chan []Pong)

	for i := 0; i < concurrentMax; i++ {
		go ping(pingChan, pongChan)
	}

	go receivePong(len(hosts), pongChan, doneChan)

	for _, ip := range hosts {
		pingChan <- ip
		// fmt.Println("sent: " + ip)
	}

	alives := <-doneChan
	// pp.Println(alives)
	f, err := os.OpenFile("allIP.txt", os.O_TRUNC|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		_ = fmt.Errorf("%v ahhah", err)
		// f, err = os.Create("allIP.txt")
		// if err != nil {
		// 	fmt.Errorf("unable to create file%v", err)
		// }
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, aliveIP := range alives {

		// fmt.Println(aliveIP.Ip)
		fmt.Fprintln(w, aliveIP.Ip)
	}
	err = w.Flush()
	if err != nil {
		fmt.Errorf("unable to flush %v", err)
	}
	// fmt.Println(alives.IP)
}
