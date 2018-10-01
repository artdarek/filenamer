package filenamer

import (
	"testing"
)

func TestFilenamer_Get(t *testing.T) {
	fnr := New("filename.jpg")
	if fnr.Get() != "filename.jpg" {
		t.Error("Expected filename.jpg")
	}
}

func TestFilenamer_CleanIt(t *testing.T) {
	fnr := New("file name.jpg")
	fnr.CleanIt()
	if fnr.Get() != "file_name.jpg" {
		t.Error("Expected file_name.jpg")
	}
}

func TestFilenamer_HashIt(t *testing.T) {
	fnr := New("filename.jpg")
	fnr.HashIt()
	if len(fnr.Get()) != 36 {
		t.Error("435ed7e9f07f740abf511a62c00eef6e.jpg")
	}
}

func TestFilenamer_SeparateWith(t *testing.T) {
	fnr := New("file name.jpg")
	fnr.SeparateWith("-")
	fnr.CleanIt()
	if fnr.Get() != "file-name.jpg" {
		t.Error("Expected file-name.jpg")
	}
}

func TestFilenamer_WithExtension(t *testing.T) {
	fnr := New("filename.jpg")
	fnr.WithExtension(".png")
	fnr.CleanIt()
	if fnr.Get() != "filename.png" {
		t.Error("Expected filename.png")
	}
}

func TestFilenamer_WithoutExtension(t *testing.T) {
	fnr := New("filename.jpg")
	fnr.WithExtensionRemoved()
	fnr.CleanIt()
	if fnr.Get() != "filename" {
		t.Error("Expected filename")
	}
}

func TestFilenamer_WithCustomPrefix(t *testing.T) {
	fnr := New("filename.jpg")
	fnr.WithCustomPrefix("test")
	fnr.CleanIt()
	if fnr.Get() != "test_filename.jpg" {
		t.Error("Expected test_filename")
	}
}

func TestFilenamer_WithCustomSuffix(t *testing.T) {
	fnr := New("filename.jpg")
	fnr.WithCustomSuffix("test")
	fnr.CleanIt()
	if fnr.Get() != "filename_test.jpg" {
		t.Error("Expected filename_test.jpg")
	}
}

func TestFilenamer_WithRandomPrefix(t *testing.T) {
	fnr := New("filename.jpg")
	fnr.WithRandomPrefix(4)
	if len(fnr.Get()) != 17 {
		t.Error("Expected rAnD_filename.jpg")
	}
}

func TestFilenamer_WithRandomSuffix(t *testing.T) {
	fnr := New("filename.jpg")
	fnr.WithRandomSuffix(4)
	if len(fnr.Get()) != 17 {
		t.Error("Expected filename_rAnD.jpg")
	}
}

func TestFilenamer_WithTimestamp(t *testing.T) {
	fnr := New("filename.jpg")
	fnr.WithTimestamp()
	if len(fnr.Get()) != 27 {
		t.Error("Expected 20181001100047_filename.jpg")
	}
}

func TestFilenamer_WithReplacement(t *testing.T) {
	fnr := New("filename.jpg")
	fnr.WithReplacement("file", "dir")
	fnr.CleanIt()
	if fnr.Get() != "dirname.jpg" {
		t.Error("Expected dirname.jpg")
	}
}