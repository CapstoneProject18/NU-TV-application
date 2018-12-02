package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "172.19.16.239:6969")
	if err != nil {
		os.Exit(1)
	}
	var buf [512]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			os.Exit(2)
		}
		line = strings.TrimRight(line, "\t\r\n")
		strs := strings.SplitN(line, " ", 2)
		if strs[0] == "reqHead" {
			conn.Write([]byte(line))
			_, err := conn.Read(buf[0:])
			if err != nil {
				fmt.Println(string(buf[0:]))
			}
		} else if strs[0] == "sm" {
			conn.Write([]byte(line))
			n, err := conn.Read(buf[0:])
			if n > 0 && err != nil {
				fmt.Println(string(buf[0:]))
			}
		}
	}

}
