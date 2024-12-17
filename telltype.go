package main

import "fmt"

func main() {
	
	fmt.Println("enter a value")
	fmt.Scanln(&i)
	var i interface{}

	switch i.(type) {
	case int:
		fmt.Println("The value is an integer")
	case string:
		fmt.Println("The value is a string")
	case float64:
		fmt.Println("The value is a float64")
	default:
		fmt.Println("The value is of an unknown type")
	}
}
