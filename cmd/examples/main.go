package main

import (
	"fmt"
	"github.com/artdarek/filenamer"
)

func main() {

	// Example 1
	fn := filenamer.New("test file.jpg")
	fn.CleanIt()
	fmt.Println(fn.Get())

	// Example 2
	fn = filenamer.New("test file.jpg")
	fn.CleanIt()
	fn.WithCustomPrefix("00001")
	fmt.Println(fn.Get())

	// Example 3
	fn = filenamer.New("test file.jpg")
	fn.CleanIt()
	fn.WithCustomPrefix("00001")
	fn.HashIt()
	fmt.Println(fn.Get())

	// Example 4
	fn = filenamer.New("test file.jpg")
	fn.CleanIt()
	fn.WithCustomPrefix("00001")
	fn.HashIt()
	fn.WithCustomSuffix("00002")
	fmt.Println(fn.Get())

	// Example 5
	fn = filenamer.New("test file.jpg")
	fn.SeparateWith("-")
	fn.CleanIt()
	fn.SeparateWith("_")
	fn.WithCustomPrefix("00001")
	fn.WithCustomSuffix("00002")
	fmt.Println(fn.Get())

	// Example 6
	fn = filenamer.New("test file.jpg")
	fn.CleanIt()
	fn.WithRandomPrefix(10)
	fmt.Println(fn.Get())

	// Example 7
	fn = filenamer.New("test file.jpg")
	fn.CleanIt()
	fn.WithRandomSuffix(5)
	fn.WithTimestamp()
	fmt.Println(fn.Get())

	// Example 8
	fn = filenamer.New("testfile.jpg")
	fn.WithExtension(".png")
	fmt.Println(fn.Get())

	// Example 9
	fn = filenamer.New("test file.jpg")
	fn.CleanIt()
	fn.WithExtensionRemoved()
	fn.WithReplacement("test","TEST")
	fmt.Println(fn.Get())
}