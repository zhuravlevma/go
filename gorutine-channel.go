package main

import "fmt"

func gorutine_channel() {
	channel1 := make(chan int, 1) // second params - buffer

	go func(in chan int) {
		val := <-in // read from channel
		fmt.Println("Value from channel", val)
		fmt.Println("after reading")
	}(channel1)

	channel1 <- 42   // put to channel
	channel1 <- 1004 // deadlock for chan without buffer
	fmt.Println("after put channel")
	fmt.Scanln()
}
