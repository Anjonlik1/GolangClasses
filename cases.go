package main

import (
	"fmt"
	"time"
)

func main() {
	var year int
	t := time.Now()
	year := t.Year()

	fmt.Println("enter a year")
	fmt.Scanln(&year)

	switch year {
	case 0:
		fmt.Println("Maymun")
		break
	case 1:
		fmt.Println("Ho'roz")
		break
	case 2:
		fmt.Println("It")
		break
	case 3:
		fmt.Println("To'ng'iz")
		break
	case 4:
		fmt.Println("Sichqon")
		break
	case 5:
		fmt.Println("Sigir")
		break
	case 6:
		fmt.Println("Yo'lbars")
		break
	case 7:
		fmt.Println("Quyon")
		break
	case 8:
		fmt.Println("dragon")
		break
	case 9:
		fmt.Println("Ilon")
		break
	case 10:
		fmt.Println("Ot")
		break
	case 11:
		fmt.Println("nimadur")
		break
	default:
		fmt.Println("Invalid day")
	}
}
