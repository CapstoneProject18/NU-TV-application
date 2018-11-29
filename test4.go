package main

import (
	"fmt"
	"net"
	"os"
	"update"
)

func main() {
	//1. dusre ke server se baat karne ke liye reqHeader ready karo
	//2. client ko abhi on karne ki koi jarurat nahi hai, kyuki uske render karwane ke liye enough data nhi hai.
	//3. ek baar apna server unse baat karne lage aur jitne user online hai, uski list magke , movies list ready kar lo
	//4. kyui agar user vo link pe click karta hai to client server ke paas request aayega aur vo phir server se mangega aur phir server tcp connection se usko file lake dega
	//5. aur aise hi client and server phir baat karte rahenge.
	//6. user kabhi bhi tcp server ko nahi chedega, vo bas http server ko requets bhejega

	time, reqh, err := update.UpdateNUtv()
	if err != nil {
		os.Exit(1)
	}
	listPeers, err := aliveServices.List(time.Now())

	tcpAddress, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	l, err := net.ListenTCP("tcp", tcpAddress)
	if err != nil {
		fmt.Errorf("unable to just listen this tcp address%v", err)
	}
	for {
		con, err := l.Accept()
		if err != nil {
			fmt.Errorf("unable to get con from this listner%v", err)
		}
		go handleCon(con)
	}
}

func handleCon(con net.Conn) {
	defer con.Close()

}
