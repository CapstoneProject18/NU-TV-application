package main

import (
	"fmt"
)

func main() {

	for {
		server()
	}

}
func server() {
	fmt.Println("inintialiasing...")

	fl := make(chan int) //file lena hai
	fd := make(chan int) //file dena hai
	go client(fd, fl)
	fd <- 3
	fmt.Println("out:=", <-fl)

}

func client(in <-chan int, out chan<- int) {

	fmt.Println("in:= ", <-in)
	out <- 6

}
