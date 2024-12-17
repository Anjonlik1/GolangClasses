package main

import (
	"fmt"
	"time"
)

func main() {

	type Employee struct {
		ID        int
		Name      string
		Address   string
		DoB       time.Time
		Position  string
		Salary    int
		ManagerID int
	}
	var dilbert Employee
	dilbert.Salary -= 5000
	fmt.Println("salary is ", dilbert.Salary)

}
