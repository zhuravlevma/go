package main

import (
	"fmt"
	"runtime"
)

func calculate(i int) {
	for j:=0; j<5; j++ {
		fmt.Println(i, j)
		runtime.Gosched()
	}
}

func GorutineRun() {
	for i:=0; i <5; i++ {
		go calculate(i)
		fmt.Println("next step")
	}
	fmt.Scanln()
}
