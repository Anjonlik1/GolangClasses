package main

import "fmt"

func calc(x int, y int) (int, int) {
	var sum1 = x + y
	var sum2 = x - y
	return sum1, sum2

}

func main() {

	num1 := 4
	num2 := 5

	result1, result2 := calc(num1, num2)
	fmt.Println(result1, result2)
}
