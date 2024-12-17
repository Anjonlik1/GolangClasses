package main

import (
	"fmt"
)

func increment(num int) {
	num++
	fmt.Println("number inside the method", num)
}
func decrement(num int) {
	num--
	fmt.Println("number inside the method", num)
}
func main() {
	num := 0
	fmt.Println("before the call, x is ", num)
	increment(num)
	fmt.Println("after the call", num)
	fmt.Println("____________________")
	fmt.Println("before the call, x is ", num)
	decrement(num)
	fmt.Println("after the call", num)

}
