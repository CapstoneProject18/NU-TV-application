package online

import (
	"bufio"
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
	fmt.Println("online list to go")
	f, err := os.Open("aliveServices.txt")
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("%v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line + "line")
		fmt.Println("asking reqHead")
		line = strings.TrimRight(line, "/n")
		go Client(line)
		// _, err = conn.Read(buf[0:])
		// if err != nil {
		// 	return fmt.Errorf("%v", err)
		// }
		// fmt.Println(string(buf[0:]))
		return nil
	}
	if err = scanner.Err(); err != nil {
		fmt.Println("i dont why?")
		return nil
	}
	return nil
}

func Client(line string) {
	conn, err := net.Dial("tcp", line+":6969")
	if err != nil {
		fmt.Println(err, "err")
		return
	}

	fw, err := os.OpenFile("othersList.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = conn.Write([]byte("reqHead rb"))
	if err != nil {
		fmt.Println(err, "errorrr")
		return
	}
	// buf := make([]byte, 512)
	fmt.Println("starting for loop")

	// decoder := gob.NewDecoder(conn)
	// fmt.Println("hello")
	// var rh ReqHeader
	// err = decoder.Decode(&rh)
	// if err != nil {
	// 	fmt.Println("error to decode rh", err)
	// 	panic(err)

	// }
	conn.Close()
	fmt.Println("hello jee")
	// OthersHeaders = append(OthersHeaders, rh)
	// fmt.Println("hello", rh.MyList[0])
	// for _, val := range rh.MyList {
	// 	fmt.Fprintf(fw, val)
	// 	fmt.Println(val)
	// }
	fmt.Println("exiting online")
	fw.Close()

}
