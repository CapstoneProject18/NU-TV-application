package server

import (
	"encoding/gob"
	"fmt"
	"net"
	"update"
)

// type ReqHeader struct {
// 	MyAddr net.Addr
// 	MyList []string
// }
// var moviePath map[string][string] //map[moviename][ movie path (ip+ file path )]
func Server(con net.Conn, reqh update.ReqHeader) {
	defer con.Close()
	fmt.Println("i m a node server which is about to serve tcp connection to all other node server in nutv network")
	encoder := gob.NewEncoder(con)
	for {
		// msgBytes, err := ioutil.ReadAll(con)

		var msgBytes [512]byte
		n, err := con.Read(msgBytes[0:])
		if err != nil {
			continue
		}

		s := string(msgBytes[0:n])
		// fmt.Print(s[0:7])
		if s[0:7] == "reqHead" {
			encoder.Encode(reqh)
			// con.Write([]byte(fmt.Sprintf("%v", reqh)))
			fmt.Println(s)
		} else if s[0:2] == "sm" {
			con.Write([]byte("ok"))
			fmt.Println("send me movie", s)
		}
	}
}
