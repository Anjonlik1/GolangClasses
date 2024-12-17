package main

import (
	"fmt"
	"time"
)

func printNumbers(n int) {
	if n <= 10 {
		fmt.Println(n)
		printNumbers(n + 1)
	}
}

func main() {
	go printNumbers(1)

	time.Sleep(time.Second * 2)
}
