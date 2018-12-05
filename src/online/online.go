package online

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
)

type ReqHeader struct {
	MyAddr net.Addr
	MyList []string
}

var OthersHeaders []ReqHeader

func OnlineList() error {
	f, err := os.Open("aliveServices.txt")
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	defer f.Close()
	fw, err := os.OpenFile("othersList.txt", os.O_TRUNC|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer fw.Close()
	scanner := bufio.NewScanner(f)
	fmt.Println("starting dailing for reqHead")
	var buf [512]byte
	for {
		_, err = f.Read(buf[0:])
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:]))
	}
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Print(line)
		fmt.Println("asking reqHead")
		line = strings.TrimRight(line, "/n")
		conn, err := net.Dial("tcp", line+":6969")
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		// var buf [512]byte
		defer conn.Close()
		conn.Write([]byte("reqHead "))
		decoder := gob.NewDecoder(conn)
		rh := &ReqHeader{}
		decoder.Decode(rh)
		// _, err = conn.Read(buf[0:])
		// if err != nil {
		// 	return fmt.Errorf("%v", err)
		// }

		OthersHeaders = append(OthersHeaders, *rh)
		for _, val := range rh.MyList {
			fmt.Fprintf(fw, val)
		}
		for _, val := range rh.MyList {
			fmt.Println(val)
		}
		// fmt.Println(string(buf[0:]))
		return nil
	}
	if err = scanner.Err(); err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}
