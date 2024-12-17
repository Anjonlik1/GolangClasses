package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}
type Author struct {
	Name  string `json: "name"`
	Age   int    `json:"author"`
	Tirik bool   `json:"is_tirik"`
}

func main() {
	// fmt.Println("hello world")
	author := Author{Name: "Pirimqul Qodirov", Age: 190, Tirik: false}
	book := Book{Title: "Yulduzli Tunlar", Author: author}

	fmt.Printf("%+v\n", book)

	byteArray, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}
