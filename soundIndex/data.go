package main

import (
	"strings"

	"github.com/xrash/smetrics"
)

type FormData struct {
	UserInput string
	Result    string
}

type Table struct {
	Name     string
	Soundex  string
	Expected string
	Test     string
}

// soundex
func soundex(input string) string {
	data := map[rune]string{
		'B': "1", 'F': "1", 'P': "1", 'V': "1",
		'C': "2", 'G': "2", 'J': "2", 'K': "2", 'Q': "2", 'S': "2", 'X': "2", 'Z': "2",
		'D': "3", 'T': "3",
		'L': "4",
		'M': "5", 'N': "5",
		'R': "6",
	}
	if len(input) == 0 {
		return ""
	}
	input = strings.ToUpper(input)
	result := string(input[0])
	var prevCode rune
	for _, val := range input[1:] {
		if code, exists := data[val]; exists {
			if val != prevCode {
				result += code
			}
			prevCode = val
		} else {
			prevCode = ' '
		}
	}
	if len(result) < 4 {
		result += strings.Repeat("0", 4-len(result))
	} else if len(result) > 4 {
		result = result[:4]
	}
	return result
}

// Основной алгоритм обработки данных
func processData(input string) []FormData {
	var data []FormData
	separator := func(r rune) bool {
		return r == ' ' || r == ','
	}
	words := strings.FieldsFunc(input, separator)
	for _, word := range words {
		data = append(data, FormData{
			UserInput: word,
			Result:    soundex(word),
		})
	}
	return data
}

// Конвертирует FormData в Table
func convertToTable(data []FormData) []Table {
	var table []Table
	for _, val := range data {
		test := "FAIL"
		if val.Result == smetrics.Soundex(val.UserInput) {
			test = "PASS"
		}
		table = append(table, Table{
			Name:     val.UserInput,
			Soundex:  val.Result,
			Expected: smetrics.Soundex(val.UserInput),
			Test:     test,
		})
	}
	return table
}
