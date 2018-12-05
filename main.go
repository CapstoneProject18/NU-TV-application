package main

import (
	"fmt"
	"log"
	"net"
	"online"
	"os"
	"os/exec"
	"runtime"
	"servehttp"
	"server"
	"strings"
	"time"
	"update"
)

func main() {
	//do u need something before to start server,,, get it here
	startTime := time.Now()
	updeTime, reqHeader, err := update.UpdateNUtv()
	if err != nil {
		fmt.Errorf("Error found in updating your NUtv app%v", err)
	}
	fmt.Println(updeTime)
	if b := startTime.Before(updeTime); b == false {
		os.Exit(33)
	}
	err = online.OnlineList()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//procedd to make server which let it run forever(app on lifespan)
	//in server only we will call client server which is will instatiate it which will also run for same life span
	//and then server and client will comminaticate with channels

	addr := strings.Split(reqHeader.MyAddr.String(), "/")
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:6969", addr[0]))
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Errorf("not able to listen%v", err)
	}
	go servehttp.ServeHttp()
	openbrowser("http://127.0.0.1:8080")
	fmt.Println("browser is just opened")
	for {
		con, err := l.Accept()
		if err != nil {
			fmt.Errorf("not able to accept listner %v", err)
		}
		fmt.Println("accpted a new connection and about to serve him")
		go server.Server(con, reqHeader)
	}
}

func openbrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

//before making connection with your own nutv client http server
//you must first check you carry updated online list
//then ask from every online node server for their req header
//you append your list to all collected list
//and start your http server with list arguments so that list can be viewed in browser
//and after you will be ready to serve your list
