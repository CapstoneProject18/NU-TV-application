package allIP

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
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
	f, err := os.OpenFile("allIP.txt", os.O_TRUNC|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		_ = fmt.Errorf("%v ahhah", err)
		// f, err = os.Create("allIP.txt")
		// if err != nil {
		// 	fmt.Errorf("unable to create file%v", err)
		// }
	}
	defer f.Close()
	for i := 0; i < pongNum; i++ {
		pong := <-pongChan
		//  fmt.Println("received:", pong)
		if pong.Alive {
			// fmt.Println("alive wow", pong)
			alives = append(alives, pong)
			fmt.Fprintln(f, pong.Ip)
		}
	}
	doneChan <- alives
	// fmt.Println(doneChan)
}
func MyAddr() net.Addr {

	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Errorf("cannot find net Interfaces")
	}
	for _, iface := range netInterfaces {
		if strings.HasPrefix(iface.Name, "wlp") {
			address, err := iface.Addrs()
			if err != nil {
				fmt.Errorf("cannot find interfaces address%v", err)
			}
			return address[0]

		}

		// fmt.Println(index, iface.Name, address)
	}
	return nil
}
func AllIP() {

	host := MyAddr()
	fmt.Print(host.String())

	hosts, _ := Hosts(host.String())
	concurrentMax := 100
	pingChan := make(chan string, concurrentMax)
	pongChan := make(chan Pong, len(hosts))
	doneChan := make(chan []Pong)

	for i := 0; i < concurrentMax; i++ {
		go ping(pingChan, pongChan)
	}

	go receivePong(len(hosts), pongChan, doneChan)

	for _, ip := range hosts {
		// fmt.Println(ip)

		pingChan <- ip
		// fmt.Println("sent: " + ip)
	}

	alives := <-doneChan
	_ = alives
	// fmt.Println(alives, "alives")
	// fmt.Fprintln(f, alives)

	// w := bufio.NewWriter(f)
	// for _, aliveIP := range alives {

	// 	fmt.Println(aliveIP.Ip)
	// 	fmt.Fprintln(w, aliveIP.Ip)
	// }
	// err = w.Flush()
	// if err != nil {
	// 	_ = fmt.Errorf("unable to flush %v", err)
	// }
	// fmt.Println(alives.IP)
}
