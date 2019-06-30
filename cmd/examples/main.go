package main

import (
	"fmt"

	"github.com/artdarek/filenamer"
)

func main() {

	// Example 1
	source := "test FiLe.jpg"
	fn := filenamer.New(source)
	fn.CleanIt()
	fmt.Println(fmt.Sprintf("%s -> %s", source, fn.Get()))

	// Example 2
	source = "test FiLe.JPG"
	fn = filenamer.New(source)
	fn.CleanIt()
	fn.WithCustomPrefix("00001")
	fmt.Println(fmt.Sprintf("%s -> %s", source, fn.Get()))

	// Example 3
	source = "test FiLe.JPG"
	fn = filenamer.New(source)
	fn.CleanIt()
	fn.WithCustomPrefix("00001")
	fn.HashIt()
	fmt.Println(fmt.Sprintf("%s -> %s", source, fn.Get()))

	// Example 4
	source = "test FiLe.JPG"
	fn = filenamer.New(source)
	fn.CleanIt()
	fn.WithCustomPrefix("00001")
	fn.HashIt()
	fn.WithCustomSuffix("00002")
	fmt.Println(fmt.Sprintf("%s -> %s", source, fn.Get()))

	// Example 5
	source = "test FiLe.JPG"
	fn = filenamer.New(source)
	fn.SeparateWith("-")
	fn.CleanIt()
	fn.SeparateWith("_")
	fn.WithCustomPrefix("00001")
	fn.WithCustomSuffix("00002")
	fmt.Println(fmt.Sprintf("%s -> %s", source, fn.Get()))

	// Example 6
	source = "test FiLe.JPG"
	fn = filenamer.New(source)
	fn.CleanIt()
	fn.WithRandomPrefix(10)
	fmt.Println(fmt.Sprintf("%s -> %s", source, fn.Get()))

	// Example 7
	source = "test FiLe.JPG"
	fn = filenamer.New(source)
	fn.CleanIt()
	fn.WithRandomSuffix(5)
	fn.WithTimestamp()
	fmt.Println(fmt.Sprintf("%s -> %s", source, fn.Get()))

	// Example 8
	source = "test FiLe.JPG"
	fn = filenamer.New(source)
	fn.WithExtension(".png")
	fmt.Println(fmt.Sprintf("%s -> %s", source, fn.Get()))

	// Example 9
	source = "test FiLe.JPG"
	fn = filenamer.New(source)
	fn.CleanIt()
	fn.WithExtensionRemoved()
	fn.WithReplacement("test","TEST")
	fmt.Println(fmt.Sprintf("%s -> %s", source, fn.Get()))

	// Example 10
	source = "test FiLe.JPG"
	fn = filenamer.New(source)
	fn.CleanIt()
	fmt.Println(fmt.Sprintf("%s -> %s", source, fn.Get()))
}