package font

import (
	"errors"
	"fmt"
)

type Font struct {
	family string
	size   int
}

func New(name string, size int) *Font {
	return &Font{family: name, size: size}
}
func (f *Font) SetFamily(family string) error {
	if family == "" {
		return errors.New("font family is empty")
	}
	f.family = family
	return nil
}
func (f *Font) SetSize(size int) error {
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
	f.size = size
	return nil
}
func (f *Font) Family() string {
	return f.family
}
func (f *Font) Size() int {
	return f.size
}
func (f *Font) String() string {
	return fmt.Sprintf(`{font-family: "%s"; font-size: %dpt;}`, f.family, f.size)
}
