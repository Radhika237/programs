package main

import (
	"fmt"
)

func pinger(c chan string, value int) {
	for i := 0; i < value; i++ {
		c <- "ping"
	}
	close(c)
}

func main() {
	var c chan string = make(chan string)
	var i int

	fmt.Print("Type a number: ")
	fmt.Scan(&i)
	fmt.Println("Your number is:", i)

	go pinger(c, i)

	opened := true
	var msg string

	for opened {
		msg, opened = <-c
		fmt.Println(msg)
	}
}
