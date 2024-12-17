package main

import "fmt"

func main() {
	arr := []int{8, 11, 15, 99, 1, 25}
	fmt.Println(findLargestNumber(arr))

}

func findLargestNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	largest := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}
	}
	return largest
}
