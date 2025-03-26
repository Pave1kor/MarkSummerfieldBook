// Copyright © 2011-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"strings"
)

func main() {
	irregularMatrix := [][]int{{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11},
		{12, 13, 14, 15},
		{16, 17, 18, 19, 20}}
	fmt.Println("irregular:", irregularMatrix)
	slice := Flatten(irregularMatrix)
	fmt.Printf("1x%d: %v\n", len(slice), slice)
	fmt.Printf(" 3x%d: %v\n", neededRows(slice, 3), Make2D(slice, 3))
	fmt.Printf(" 4x%d: %v\n", neededRows(slice, 4), Make2D(slice, 4))
	fmt.Printf(" 5x%d: %v\n", neededRows(slice, 5), Make2D(slice, 5))
	fmt.Printf(" 6x%d: %v\n", neededRows(slice, 6), Make2D(slice, 6))
	slice = []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	fmt.Println("Original:", slice)
	slice = UniqueInts(slice)
	fmt.Println("Unique:  ", slice)

	iniData := []string{
		"; Cut down copy of Mozilla application.ini file",
		"",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0",
		"EnableExtensionManager=1",
	}
	ini := ParseIni(iniData)
	PrintIni(ini)
}
func UniqueInts(array []int) []int {
	unique := make([]int, 0, len(array))
	seen := map[int]bool{}
	for _, x := range array {
		if !seen[x] {
			unique = append(unique, x)
			seen[x] = true
		}
	}
	return unique
}
func Flatten(matrix [][]int) []int {
	slice := make([]int, 0, len(matrix)+len(matrix[0]))
	for _, innerSlice := range matrix {
		slice = append(slice, innerSlice...)
	}
	return slice
}
func Make2D(slice []int, columns int) [][]int {
	matrix := make([][]int, neededRows(slice, columns))
	for i, x := range slice {
		row := i / columns
		column := i % columns
		if matrix[row] == nil {
			matrix[row] = make([]int, columns)
		}
		matrix[row][column] = x
	}
	return matrix
}

func neededRows(slice []int, columns int) int {
	rows := len(slice) / columns
	if len(slice)%columns != 0 {
		rows++
	}
	return rows
}

func ParseIni(iniData []string) map[string]map[string]string {
	foreignMap := make(map[string]map[string]string)
	var section string
	for _, line := range iniData {
		switch {
		case len(line) == 0:
			continue
		case line[0] == ';':
			continue
		case line[0] == '[' && line[len(line)-1] == ']':
			section = line[1 : len(line)-1]
			foreignMap[section] = make(map[string]string)
		default:
			parts := strings.SplitN(line, "=", 2)
			foreignMap[section][parts[0]] = parts[1]
		}
	}
	return foreignMap
}

func PrintIni(ini map[string]map[string]string) {
	for section, values := range ini {
		fmt.Println("[" + section + "]")
		for key, value := range values {
			fmt.Println(key + "=" + value)
		}
	}
}
