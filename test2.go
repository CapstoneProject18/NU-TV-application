package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"servehttp"
	"strings"
	"time"
	"update"
)

type decIden struct {
	ipAddr net.TCPAddr
}

var globalIden []decIden

func main() {
	updeTime, reqHeader, err := update.UpdateNUtv()
	if err != nil {
		fmt.Errorf("Error found in updating your NUtv app%v", err)
	}
	go servehttp.ServeHttp()
	_ = updeTime
	fmt.Println(reqHeader.MyAddr.String())
	// var dataChannel chan string
	addr := strings.Split(reqHeader.MyAddr.String(), "/")
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:6969", addr[0]))
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Errorf("not able to listen%v", err)
	}
	for {
		con, err := l.Accept()
		if err != nil {
			fmt.Errorf("not able to accept listner %v", err)
		}
		_ = con

		// serve(dataChannel)
		fmt.Print("i m going to b")
		go handle(con)

	}
}
func handle(con net.Conn) {
	fmt.Println(con)
	go fetch()

}

func fetch() {
	if confirmAllIp() {
		f, err := os.Open("allIP.txt")
		if err != nil {
			fmt.Errorf("unabe to open allIP.txt %v", err)
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			v := scanner.Text()
			tcpAddr, err := net.ResolveTCPAddr("tcp", v)
			if err != nil {
				fmt.Errorf("dont know why?? %v", err)
			}
			fmt.Println("i m going to die")
			cmdV := exec.Command(fmt.Sprintf("curl %s:6969", tcpAddr))
			time.Sleep(time.Millisecond * 9)
			ret := exec.Command(fmt.Sprintf("^c"))
			_ = ret
			fmt.Println(cmdV)
		}
	}

}

func confirmAllIp() bool {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Errorf("unable to read dir to find to allIP.txt %v", err)
	}
	for _, file := range files {
		if file.Name() == "allIP.txt" {
			return true
		}
	}
	return false
}
