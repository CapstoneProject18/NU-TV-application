package server

import (
	"bufio"
	"fmt"
	"net"
	"update"
)

// type ReqHeader struct {
// 	MyAddr net.Addr
// 	MyList []string
// }
func Server(con net.Conn, reqh update.ReqHeader) {
	defer con.Close()
	// scanner := bufio.NewScanner(con)
	// for scanner.Scan() {
	// 	v := scanner.Text()
	// 	// switch v {
	// 	// case "reqh":
	// 	// 	{

	// 	// 	}
	// 	// case "":
	// 	// 	{

	// 	// 	}
	// 	// }
	// 	fmt.Println(v)
	// }
	// if scanner.Err() != nil {
	// 	fmt.Errorf("out time ")
	// }
	reader := bufio.NewReader(con)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			fmt.Errorf("error from reader%v", err)
		}
		fmt.Print(string(b))
		con.Write([]byte("sk"))
	}

}

//yaha start hota hai jab koi new peer apna aplication start karta hai, aur apse connection leta hai, ye depend karta hai
//aap kya serve karte hai
//jo jo aliveServices.txt me pinned hai, unse connect karke unka req header accpet karo, list banao aur
//ye bhi possiblity hai ki koi peers baad me aake apna req header de , to usko bhi append karna hoga

// func client() {

// }
