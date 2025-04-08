package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	arr := make([]string, 0)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	fmt.Print("prefix: " + CommonPrefix(arr))

}
func CommonPrefix(arr []string) string {
	var result strings.Builder
	if len(arr) == 0 {
		return ""
	}
	for j := 0; j < len(arr[0]); j++ {
		result.WriteByte(arr[0][j])
		for i := 1; i < len(arr); i++ {
			if !strings.HasPrefix(arr[i], result.String()) {
				return result.String()[0:j]
			}
		}
	}
	// Если все строки идентичны
	return arr[0]
}
