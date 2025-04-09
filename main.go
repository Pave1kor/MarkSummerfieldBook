package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

type File struct {
	current string
	next    string
}

func main() {
	testData := [][]string{
		{"/home/user/goeg", "/home/user/goeg/prefix",
			"/home/user/goeg/prefix/extra"},
		{"/home/user/goeg", "/home/user/goeg/prefix",
			"/home/user/prefix/extra"},
		{"/pecan/π/goeg", "/pecan/π/goeg/prefix",
			"/pecan/π/prefix/extra"},
		{"/pecan/π/circle", "/pecan/π/circle/prefix",
			"/pecan/π/circle/prefix/extra"},
		{"/home/user/goeg", "/home/users/goeg",
			"/home/userspace/goeg"},
		{"/home/user/goeg", "/tmp/user", "/var/log"},
		{"/home/mark/goeg", "/home/user/goeg"},
		{"home/user/goeg", "/tmp/user", "/var/log"},
	}

	for _, data := range testData {
		fmt.Printf("commonPrefix(%q, %q, %q) = %s\n",
			data[0], data[1], data[2], commonPrefix(data[0], data[1], data[2]))
	}
}
func commonPrefix(str ...string) string {
	var result [][]string
	res := strings.Split(filepath.Clean(str[0]), string(filepath.Separator))
	cols := len(res)
	for _, value := range str[1:] {
		res := strings.Split(filepath.Clean(value), string(filepath.Separator))
		if len(res) < cols {
			cols = len(res)
		}
		result = append(result, res)
	}
	rows := len(result)
	common := 1
	for col := 1; col < cols; col++ {
		for row := 1; row < rows; row++ {
			if result[row][col] != result[0][col] {
				return strings.Join(result[row][:col], string(filepath.Separator))
			}
		}
		common++
	}
	return strings.Join(result[0][:common], string(filepath.Separator))
}
