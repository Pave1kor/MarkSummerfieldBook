package main

import (
	"reflect"
	"testing"
)

func TestArchiveFileList(t *testing.T) {
	file := "archive.tar.bz2"
	expected := []string{"archive.tar"} // Replace with the expected file names
	files, err := ArchiveFileList(file)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if !reflect.DeepEqual(files, expected) {
		t.Errorf("Expected %v, but got %v", expected, files)
	}
}
