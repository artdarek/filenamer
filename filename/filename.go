package filename

import (
	"path/filepath"
	"strings"
)

// Filename stores file name and extension
type Filename struct {
	name      string // file name
	extension string // File extension
}

// New creates and returns struct that keeps
// file name and extension from given filename
func New(filename string) Filename {
	name, extension := getNameAndExtension(filename)
	return Filename{name: name, extension: extension}
}

// Get returns filename created from name and extension
func (f *Filename) Get() string {
	return f.name + f.extension
}

// GetName returns file name
func (f *Filename) GetName() string {
	return f.name
}

// SetName sets name for a file
func (f *Filename) SetName(name string) {
	f.name = name
}

// GetExtension returns file extension
func (f *Filename) GetExtension() string {
	return f.extension
}

// SetExtension sets file extensio
func (f *Filename) SetExtension(extension string) {
	f.extension = extension
}

// SetPrefix sets filename prefix with given separator
func (f *Filename) SetPrefix(text string, separator string) {
	f.name = text + separator + f.GetName()
}

// SetSuffix sets filename suffix with given separator
func (f *Filename) SetSuffix(text string, separator string) {
	f.name = f.GetName() + separator + text
}

// getNameAndExtension returns name and extension for given filename as separate strings
func getNameAndExtension(filename string) (string, string) {
	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	return name, ext
}
