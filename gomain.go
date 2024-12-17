package main

import (
	"fmt"
	"time"
)

func printNum(num string) {
	fmt.Println(num)
}

func main() {
	go printNum("1")
	go printNum("2")
	go printNum("3")
	go printNum("4")
	go printNum("5")
	go printNum("6")
	go printNum("7")
	go printNum("8")
	go printNum("9")
	go printNum("10")

	time.Sleep(time.Second * 2)

	fmt.Println("Hello")

}
