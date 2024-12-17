package main

import "fmt"

func main() {
	nums := []int{1, 2, 4, 5, 6}
	missing := findMissing(nums)
	fmt.Println("missing number", missing)

}

func findMissing(nums []int) int {
	n := len(nums) + 1
	sum := n * (n + 1) / 2

	for _, num := range nums {
		sum = sum - num

	}
	return sum
}
