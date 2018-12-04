package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "172.19.23.8:6969")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("reqHead yes"))
	if err != nil {
		fmt.Println(err)
	}
}
