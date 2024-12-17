package main

import (
	"fmt"
	"sync"
)

func numberPrinting(num int, wg *sync.WaitGroup) {

	fmt.Printf("Number %d printing \n", num)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go numberPrinting(i, &wg)
	}
	wg.Wait()

	fmt.Println("Finished")
}
