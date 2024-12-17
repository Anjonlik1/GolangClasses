package main

import (
	"fmt"
)

func main() {
	var value interface{} = 2 // This could be an int or a string

	if num, ok := value.(int); ok {

		fmt.Println("Value is an int:", num)
	} else if str, ok := value.(string); ok {

		fmt.Println("Value is a string:", str)
	} else if rn, ok := value.(rune); ok {

		fmt.Println("Value is a string:", rn)
	} else {

		fmt.Println("Value is of an unexpected type:", value)
	}
}
