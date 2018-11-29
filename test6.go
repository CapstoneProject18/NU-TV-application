package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// var serviceChan chan string = make(chan, string)
	f, err := os.Open("allIP.txt")

	if err != nil {
		fmt.Errorf("not able to read allIP.txt %v", err)
	}
	defer f.Close()
	err = os.Remove("aliveServices.txt")
	if err != nil {
		fmt.Errorf("not able to delete aliveServices.txt %v", err)
	}
	fw, err := os.Create("aliveServices.txt")
	if err != nil {
		fmt.Errorf("not able to create aliveServices.txt %v", err)
	}
	defer fw.Close()
	in := make(chan string, 1)
	out := make(chan string, 1)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		aliveIP := scanner.Text()
		in <- aliveIP
		go runningService(in, out)
		select {
		case v := <-out:
			{
				fmt.Println(v)
				fmt.Fprintln(fw, v)
			}
		default:
			break
		}
	}
	if scanner.Err() != nil {
		fmt.Errorf("error from scanner%v", err)
		return
	}
	// f, err := os.OpenFile

}

func runningService(in <-chan string, out chan<- string) {
	fmt.Println("i m being called")
	// err := os.Remove("aliveServices.txt")
	// if err != nil {
	// 	fmt.Errorf("unable to remove older data of aliveServices.txt%v", err)
	// }
	// f, err := os.Create("aliveServices.txt")
	// if err != nil {
	// 	fmt.Errorf("unable to open aliveServices.txt %v", err)
	// 	return
	// }
	// w := bufio.NewWriter(f)
	// defer f.Close()
	// tcpAddress, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:6969", aliveIP))
	// if err != nil {
	// 	fmt.Errorf("big mastakes in allIP files might %v", err)
	// }
	aliveIP := <-in
	con, err := net.Dial("tcp", fmt.Sprintf("%s:6969", aliveIP))

	if err != nil {
		fmt.Errorf("cannot establish connection to this ip %v", err)
		//out <- nil
		return
	}
	defer con.Close()
	fmt.Println("hurrey, connection established", con)
	fmt.Println(aliveIP)
	out <- aliveIP
	// f.WriteString(fmt.Sprintf("%s\n", aliveIP))
	// _, err = fmt.Fprintln(w, aliveIP)
	// if err != nil {
	// 	fmt.Errorf("unabe to write to aliveServices.txt %v", err)
	// }
	// err = w.Flush()
	// if err != nil {
	// 	fmt.Errorf("unable to flush %v", err)
	// }

}
