package main

import "fmt"

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
		"timothy": 30,
		"andres":  40,
	}

	fmt.Println(ages)
	fmt.Println(ages["timothy"])
}
