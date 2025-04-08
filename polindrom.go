package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	fmt.Print(IsPalindrome(str))
}

func IsPalindrome(str string) bool {
	// Очистка строки: оставляем только буквы/цифры и приводим к нижнему регистру
	var cleaned strings.Builder
	for _, r := range str {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned.WriteRune(unicode.ToLower(r))
		}
	}
	s := cleaned.String()

	// Проверка на палиндром
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-i-1] {
			return false
		}
	}
	return true
}
