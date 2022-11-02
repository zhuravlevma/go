package main

import "fmt"

func gorutine_channel2() {
	channel1 := make(chan int)
	go func(ch1 chan int) {
		for i := 0; i <= 5; i++ {
			fmt.Println("before", i)
			ch1 <- i
			fmt.Println("after", i)
		}
		close(ch1)
		print("stage finished")
	}(channel1)

	for i := range channel1 {
		fmt.Println("\tget", i)
	}
}
