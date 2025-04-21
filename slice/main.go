package main

import (
	"fmt"

	slice "slicecode/slice/slice"
)

// This function is the main function of the program
func main() {
	comparable := func(a, b interface{}) bool {
		return a.(float64) <= b.(float64)
	}

	slice := slice.New(comparable)
	slice.Add(2.2)
	slice.Add(1.1)
	slice.Add(3.3)
	slice.Add(4.4)
	slice.Len()
	fmt.Print(slice)
}
