package main

import (
	"fmt"

	fnt "main.go/font"
)

func main() {
	font := fnt.New("Arial", 12)
	err := font.SetFamily("Arial")
	if err != nil {
		fmt.Println(err)
	}
	err = font.SetSize(12)
	if err != nil {
		fmt.Println(err)
	}
	size := font.Size()
	fmt.Println(size)
	family := font.Family()
	fmt.Println(family)
	fmt.Println(font)
}
