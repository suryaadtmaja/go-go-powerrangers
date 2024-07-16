package main

import "fmt"

func readChannel(c <-chan int) {
	for {
		msg := <-c
		fmt.Println(msg, "wtf")
	}
}

func goBuffered() {
	bufferedChannel := make(chan int, 2)
	bufferedChannel <- 23
	bufferedChannel <- 24
	go readChannel(bufferedChannel)

	var input string
	fmt.Scanln(&input)
}
