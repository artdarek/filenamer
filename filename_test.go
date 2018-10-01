package filenamer

import (
	"fmt"
	"testing"
)

func TestFilename_Get(t *testing.T) {
	fn := NewFilename("filename.jpg")
	if fn.Get() != "filename.jpg" {
		t.Error("Expected filename.jpg")
	}
}

func TestFilename_GetName(t *testing.T) {
	fn := NewFilename("filename.jpg")
	if fn.GetName() != "filename" {
		t.Error("Expected filename")
	}
}

func TestFilename_GetExtension(t *testing.T) {
	fn := NewFilename("filename.jpg")
	if fn.GetExtension() != ".jpg" {
		t.Error("Expected .jpg")
	}
}

func TestFilename_SetExtension(t *testing.T) {
	fn := NewFilename("filename.jpg")
	fn.SetExtension(".gif")
	if fn.Get() != "filename.gif" {
		t.Error("Expected filename.gif")
	}
}

func TestFilename_SetName(t *testing.T) {
	fn := NewFilename("filename.jpg")
	fn.SetName("newname")
	if fn.Get() != "newname.jpg" {
		t.Error("Expected newname.jpg")
	}
}

func TestFilename_SetPrefix(t *testing.T) {
	fn := NewFilename("filename.jpg")
	fn.SetPrefix("test", "-")
	if fn.Get() != "test-filename.jpg" {
		t.Error("Expected test-filename.jpg")
	}
}

func TestFilename_SetSuffix(t *testing.T) {
	fn := NewFilename("filename.jpg")
	fn.SetSuffix("test", "-")
	fmt.Print(fn.Get())
	if fn.Get() != "filename-test.jpg" {
		t.Error("Expected filename-test.jpg")
	}
}