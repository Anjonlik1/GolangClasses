package main

import "fmt"

func main() {

	// define the variables we want to Swap
	var number1, number2, temp int

	// initializing the variables
	number1 = 45
	number2 = 63

	// printing the numbers before swapping
	fmt.Println("Numbers before swapping: \n Number 1 =", number1, "\n Number 2 =", number2)

	// swapping the numbers
	temp = number1
	number1 = number2
	number2 = temp

	// printing the numbers after swapping
	fmt.Println("Numbers after swapping:\n Number 1 =", number1, "\n Number 2 =", number2, "\n(Swap within the function)")
}
