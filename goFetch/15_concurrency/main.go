package main

/**
Channels are pipes that allow you to have your go routines communicate with each other
*/

import (
	"fmt"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)

	// // go routine
	// go func() {
	// 	count("sheep")
	// 	wg.Done()
	// }()

	// wg.Wait()

	// Run the code in receiveChannels.go
	// receive.Sel()

	c := make(chan string)
	go count("sheep", c)

	// Receiving a value from the channel
	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 6; i++ {
		fmt.Println(i, thing)

		// Sending to the Channel
		c <- thing

		time.Sleep(time.Millisecond * 400)
	}
	// Once the for loop is done sending a value the channel will be closes
	// this is useful so that you do not get the deadlock problem
	close(c)
}
