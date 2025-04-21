package main

import (
	"fmt"

	slice "slicecode/slice/slice"
)

// This function is the main function of the program
func main() {
	comparable := func(a, b interface{}) bool {
		return a.(int) <= b.(int)
	}

	slice := slice.New(comparable)
	for _, x := range []int{5, 8, -1, 3, 4, 22} {
		slice.Add(x)
	}
	fmt.Println(slice)
	for _, x := range []int{5, 5, 6} {
		slice.Add(x)
	}
	slice.Remove(5)
	value := slice.CheckList([]int{-1, 3, 4, 5, 5, 6, 8, 22})
	fmt.Printf("%v\n", value)
	fmt.Println(slice)
	slice.Len()
}
