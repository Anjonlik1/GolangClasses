package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	fmt.Println("hello world")

	book := Book{Title: "Yulduzli Tunlar", Author: "Primqul Qodirov"}

	fmt.Printf("%+v\n", book)

	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}
