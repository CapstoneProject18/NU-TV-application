package main

import (
	"fmt"
	"net"
	"server"
	"strings"
	"update"
)

func main() {
	//do u need something before to start server,,, get it here
	updeTime, reqHeader, err := update.UpdateNUtv()
	if err != nil {
		fmt.Errorf("Error found in updating your NUtv app%v", err)
	}
	_ = updeTime

	//procedd to make server which let it run forever(app on lifespan)
	//in server only we will call client server which is will instatiate it which will also run for same life span
	//and then server and client will comminaticate with channels

	addr := strings.Split(reqHeader.MyAddr.String(), "/")
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:6969", addr[0]))
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Errorf("not able to listen%v", err)
	}
	for {
		con, err := l.Accept()
		// signalin := make(chan int)
		// signalout := make(chan int)
		if err != nil {
			fmt.Errorf("not able to accept listner %v", err)
		}
		fmt.Println("accpted a new connection and about to serve him")
		go server.Server(con, reqHeader)
	}
}

//before making connection with your own nutv client http server
//you must first check you carry updated online list
//then ask from every online node server for their req header
//you append your list to all collected list
//and start your http server with list arguments so that list can be viewed in browser
//and after you will be ready to serve your list
