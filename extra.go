package main

import "fmt"

func main() {
	// bytes
	var b byte
	fmt.Printf("The zero value of b is %v, and its type is %T\n", b, b)

	b = 'b'
	fmt.Printf("The value of b is %v, and its type is %T\n", b, b)
	fmt.Printf("The character representation of b is %c, and its type is %T\n", b, b)

	// Assignment
	// Declare a variable c to be of type byte
	// then assign a value of 99 to it
	// print decimal value, character representation of it
	// and type using the fmt.Printf function.
	// Output should look like this:
	// The value of c is 99.
	// The character representation is c.
	// The type is uint8.

	// runes
	var a rune

	fmt.Printf("The zero value of a is %v, and its type is %T\n", a, a)

	a = '🦊'

	fmt.Printf("The character is %c\n", a)
	fmt.Printf("The type of a is %T\n", a)
	fmt.Printf("The unicode value is %U\n", a)

	// r := 'ᚱ'
	// r := '\u16b1'
	// fmt.Printf("The character is %c\n", r)
	// fmt.Printf("The type of r is %T\n", r)
	// fmt.Printf("The unicode value is %U\n", r)

	// var a byte = 'a'
	// b := 98
	// c := 'c'

	// fmt.Printf("Type of a is %T\n", a)
	// fmt.Printf("Type of b is %T\n", b)
	// fmt.Printf("Type of c is %T\n", c)

	// Assignment
	// Visit this website: https://unicode-table.com/en/
	// Choose any letter, symbol or emoji that you like
	// and use it to create new rune variable.
	// Then print the variable character representation and
	// also its Unicode value.

}
