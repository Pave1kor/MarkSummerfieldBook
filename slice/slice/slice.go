package slice

import (
	"fmt"
	"strings"
)

type Slicer interface {
	Clear()
	Add(value any)
	Remove(index int) bool
	Index(value any) int
	At(index int) any
	Len() int
	fmt.Stringer
}

type Slice struct {
	data    []interface{}
	compare func(a, b interface{}) bool
}

func New(compare func(a, b interface{}) bool) *Slice {
	return &Slice{
		compare: compare,
		data:    make([]interface{}, 0),
	}
}

func (s *Slice) Clear() {
	s.data = make([]interface{}, 0)
}

func (s *Slice) Add(value interface{}) {
	index := 0
	for i, v := range s.data {
		if s.compare(v, value) {
			index = i + 1
		} else {
			break
		}
	}
	s.data = append(s.data[:index], append([]interface{}{value}, s.data[index:]...)...)
}

func (s *Slice) Remove(index int) bool {
	if index < 0 || index >= len(s.data) {
		return false
	}
	s.data = append(s.data[:index], s.data[index+1:]...)
	return true
}

func (s *Slice) Index(value interface{}) int {
	for i, v := range s.data {
		if v == value { // Простое сравнение через ==
			return i
		}
	}
	return -1
}

func (s *Slice) At(index int) interface{} {
	return s.data[index]
}

func (s *Slice) Len() int {
	return len(s.data)
}

func (s *Slice) String() string {
	if s == nil || s.data == nil {
		return "data: []"
	}

	var builder strings.Builder
	builder.WriteString("data: [")

	for i, v := range s.data {
		if i > 0 {
			builder.WriteString(" ")
		}
		fmt.Fprint(&builder, v)
	}

	builder.WriteString("]")
	return builder.String()
}

// Специализированные конструкторы для конкретных типов
func NewStringSlice() *Slice {
	compare := func(a, b interface{}) bool {
		return a.(string) < b.(string)
	}
	return &Slice{
		compare: func(a, b interface{}) bool {
			return compare(a.(string), b.(string))
		},
		data: make([]interface{}, 0),
	}
}

func NewIntSlice() *Slice {
	compare := func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	return &Slice{
		compare: func(a, b interface{}) bool {
			return compare(a.(int), b.(int))
		},
		data: make([]interface{}, 0),
	}
}
