package aliveServices

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//GetRunningServices error write result to aliveservices.txt
func GetRunningServices() error {
	// var serviceChan chan string = make(chan, string)
	f, err := os.Open("allIP.txt")

	if err != nil {
		return fmt.Errorf("not able to read allIP.txt %v", err)
	}
	defer f.Close()
	err = os.Remove("aliveServices.txt")
	if err != nil {
		return fmt.Errorf("not able to delete aliveServices.txt %v", err)
	}
	fw, err := os.Create("aliveServices.txt")
	if err != nil {
		return fmt.Errorf("not able to create aliveServices.txt %v", err)
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
	con, err := net.Dial("tcp", fmt.Sprintf("%s:6969", aliveIP))
	if err != nil {
		_ = fmt.Errorf("cannot establish connection to this ip %v", err)
		return
	}
	defer con.Close()
	fmt.Println("hurrey, connection established", con)
	fmt.Println(aliveIP)
	out <- aliveIP
	fmt.Println(aliveIP)
}
