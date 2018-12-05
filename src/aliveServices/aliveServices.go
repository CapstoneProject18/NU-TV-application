package aliveServices

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//GetRunningServices error write result to aliveservices.txt
func GetRunningServices() error {
	// var serviceChan chan string = make(chan, string)
	f, err := os.Open("allIP.txt")

	if err != nil {
		return fmt.Errorf("not able to read allIP.txt %v", err)
	}
	defer f.Close()

	fw, err := os.OpenFile("aliveServices.txt", os.O_TRUNC|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer fw.Close()
	in := make(chan string, 1)
	out := make(chan string, 1)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		aliveIP := scanner.Text()
		// fmt.Println(aliveIP)
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
		return fmt.Errorf("error from scanner%v", err)

	}
	return nil
}

func runningService(in <-chan string, out chan<- string) {
	aliveIP := <-in
	aliveIP = strings.TrimRight(aliveIP, "\n")

	con, err := net.Dial("tcp", fmt.Sprintf("%s:6969", aliveIP))
	if err != nil {
		_ = fmt.Errorf("cannot establish connection to this ip %v", err)
		return
	}
	con.Close()
	// _, _ = fw.WriteString(aliveIP)

	fmt.Println("hurrey, connection established with", aliveIP)
	out <- aliveIP

}
