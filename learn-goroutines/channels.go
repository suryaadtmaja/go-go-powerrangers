package main

import (
	"fmt"
	"time"
)

// Bidirectional Channels: These channels can be used for both sending and receiving.
// var c chan int
// c = make(chan int)

// Send-Only Channels: These channels can only be used for sending values.
// var sendOnly chan<- int
// sendOnly = make(chan int)

// Receive-Only Channels: These channels can only be used for receiving values.
// var receiveOnly <-chan int
// receiveOnly = make(chan int)

// The pinger function can only send to the channel
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping pong"
	}
}

// The printer function can only receive from the channel
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg, "wtf")
		time.Sleep(time.Second * 1)
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func mainChannel() {
	// Create a bidirectional channel
	var c chan string = make(chan string)

	// Pass the channel to pinger and printer with direction constraints
	go pinger(c)
	go printer(c)
	go ponger(c)

	var input string
	fmt.Scanln(&input)
}

// challenge
// How do you specify the direction of a channel type?

// by using <-chan <-chan or  you can specify the direction of channel
// chan<- can only send the channel
// <-chan can only receive
