package main

import (
	"fmt"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("ping%d", i)
	}
}

func printer(c chan string) {
	for {
		var cout chan int = make(chan int)
		msg := <-c
		fmt.Println(msg)
		go count(cout)

		for {
			// if msg := <-cout; len(msg) == 0 {
			// 	break
			// }
			r, ok := <-cout
			if r == 0 && ok == false {
				break
			}

		}

	}
}
func count(cout chan int) {
	for i := 0; i < 9; i++ {
		fmt.Println(i)
		cout <- i
	}
	close(cout)

}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
