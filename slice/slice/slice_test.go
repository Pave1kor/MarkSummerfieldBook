package slice

import (
	"fmt"
	"testing"
)

func TestSortedIntList(t *testing.T) {
	slice := NewIntSlice() // Используем ваш конструктор
	for _, x := range []int{5, 8, -1, 3, 4, 22} {
		slice.Add(x)
	}
	checkList(slice, []int{-1, 3, 4, 5, 8, 22}, t)

	for _, x := range []int{5, 5, 6} {
		slice.Add(x)
	}
	checkList(slice, []int{-1, 3, 4, 5, 5, 5, 6, 8, 22}, t)

	if slice.Index(4) != 2 {
		t.Fatal("4 missing")
	}

	printSlice(slice)

	if slice.Index(99) != -1 {
		t.Fatal("99 wrongly present")
	}

	if slice.Remove(99) != false {
		t.Fatal("99 wrongly removed")
	}

	if slice.Remove(5) != true { // Удаляем первое вхождение 5
		t.Fatal("5 not removed")
	}
	checkList(slice, []int{-1, 3, 4, 5, 5, 6, 8, 22}, t)

	if slice.Remove(5) != true {
		t.Fatal("5 not removed")
	}
	checkList(slice, []int{-1, 3, 4, 5, 6, 8, 22}, t)

	if slice.Remove(5) != true {
		t.Fatal("5 not removed")
	}
	checkList(slice, []int{-1, 3, 4, 6, 8, 22}, t)

	if slice.Index(5) != -1 {
		t.Fatalf("5 wrongly present at %d", slice.Index(5))
	}

	printSlice(slice)
	slice.Clear()

	if slice.Len() != 0 {
		t.Fatal("cleared list isn't empty")
	}

	if slice.Remove(9) != false {
		t.Fatal("9 wrongly removed")
	}

	if slice.Index(9) != -1 {
		t.Fatal("9 wrongly found")
	}
}

func printSlice(slice *Slice) {
	fmt.Print("[")
	sep := ""
	for i := 0; i < slice.Len(); i++ {
		fmt.Print(sep, slice.At(i))
		sep = ", "
	}
	fmt.Println("]")
}

func checkList(slice *Slice, ints []int, t *testing.T) {
	if slice.Len() != len(ints) {
		t.Fatalf("slice.Len()=%d != %d", slice.Len(), len(ints))
	}

	for i := 0; i < slice.Len(); i++ {
		val := slice.At(i)
		intVal, ok := val.(int)
		if !ok {
			t.Fatalf("value at %d is not int: %v", i, val)
		}
		if intVal != ints[i] {
			t.Fatalf("mismatch At(%d) %v vs. %d", i, intVal, ints[i])
		}
	}
}
