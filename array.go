package main

import "fmt"

func main() {
	var a = [7]int{3, 4, 5, 6, 5, 8, 9}
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf(" %d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf(" %d", v)
	}
}
