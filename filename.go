package filenamer

import (
	"path/filepath"
	"strings"
)

type Filename struct {
	name string
	extension string
}

func NewFilename(filename string) Filename {
	name, extension := getNameAndExtension(filename)
	return Filename{name: name, extension: extension}
}

func (f *Filename) Get() string {
	return f.name + f.extension
}

func (f *Filename) GetName() string {
	return f.name
}

func (f *Filename) SetName(name string) {
	f.name = name
}

func (f *Filename) GetExtension() string {
	return f.extension
}

func (f *Filename) SetExtension(extension string) {
	f.extension = extension
}

// Set string as prefix with given separator
func (f *Filename) SetPrefix(text string, separator string) {
	f.name = text + separator + f.GetName()
}

// Set string as suffix with given separator
func (f *Filename) SetSuffix(text string, separator string) {
	f.name = f.GetName() + separator + text
}

// Returns name and extension for given filename as separate strings
func getNameAndExtension(filename string) (string, string) {
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	return name, ext
}
