package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	// Create a new scanner that reads from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt the user for input
	fmt.Print("Enter your value: ")

	// Read the next line of input
	scanner.Scan()

	// Get the input as a string
	var name = scanner.Text()

	// Print a greeting
	fmt.Println("Kiritilgan ma'lumot %s tipida", name)
	fmt.Println("default value of %v:  ", reflect.TypeOf(name))

	whatAmI := func(name interface{}) {
		switch t := name.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
