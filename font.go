package font

import (
	"errors"
	"fmt"
)

type Font struct {
	name string
	size int
}

func NewFont() *Font {
	return &Font{name: "Arial", size: 14}
}
func (f *Font) Write(name string, size int) error {
	if name == "" {
		f.name = "default"
		return errors.New("font name is empty, set font name = default")
	}
	if size < 5 && size > 144 {
		switch size < 5 {
		case true:
			f.size = 5 // set default font size
			return errors.New("font size is too small, set font size = 5")
		case false:
			f.size = 144 // set default font size
			return errors.New("font size is not valid, set font size = 144")
		}
	}
	f.name = name
	f.size = size
	return nil
}
func (f *Font) Read() (string, int) {
	return f.name, f.size
}
func (f *Font) String() string {
	return fmt.Sprintf("Font name: %s, Font size: %d", f.name, f.size)
}
func main() {
	font := NewFont()
	err := font.Write("Arial", 12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(font)
	str, size := font.Read()
	fmt.Println(str, size)
}
