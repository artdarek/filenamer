package filenamer

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"
)

type Filenamer struct {
	input  Filename
	output Filename
	separator string
}

// Create new Filenamer instance
func New(filename string) Filenamer {
	input := NewFilename(filename)
	output := NewFilename(filename)
	return Filenamer{input: input, output: output, separator: "_"}
}

// Get output filename
func (fn *Filenamer) Get() string {
	return fn.output.Get()
}

// Add timestamp to filename
func (fn *Filenamer) WithTimestamp() {
	fn.output.SetPrefix(getFormatedTimestamp(), fn.separator)
	fn.output.SetName(fn.output.GetName())
}

// Add unique random string to filename
func (fn *Filenamer) WithRandomPrefix(length int) {
	fn.output.SetPrefix(getRandomStringWithCharset(length), fn.separator)
	fn.output.SetName(fn.output.GetName())
}

// Add unique random string to filename
func (fn *Filenamer) WithRandomSuffix(length int) {
	fn.output.SetSuffix(getRandomStringWithCharset(length), fn.separator)
	fn.output.SetName(fn.output.GetName())
}

// Add custom prefix to filename
func (fn *Filenamer) WithCustomPrefix(prefix string) {
	fn.output.SetPrefix(prefix, fn.separator)
	fn.output.SetName(fn.output.GetName())
}

// Add custom suffix to filename
func (fn *Filenamer) WithCustomSuffix(suffix string) {
	fn.output.SetSuffix(suffix, fn.separator)
	fn.output.SetName(fn.output.GetName())
}

// Add new extension
func (fn *Filenamer) WithExtension(extension string) {
	fn.output.SetExtension( extension )
}

// Remove new extension
func (fn *Filenamer) WithExtensionRemoved() {
	fn.output.SetExtension("")
}

// Replace part of a name with different string
func (fn *Filenamer) WithReplacement(current string, target string) {
	f := strings.Replace(fn.output.GetName(), current,  target, -1)
	fn.output.SetName(f)
}

// Set active separator for all feature filename manipulations
func (fn *Filenamer) SeparateWith(separator string) {
	fn.separator = separator
}

// Clean all undesired characters from filename
func (fn *Filenamer) CleanIt() {
	fn.output.SetName( strings.Replace(fn.output.GetName(), " ",  fn.separator, -1) )
	fn.output.SetName( strings.ToLower(fn.output.GetName()) )
	fn.output.SetExtension( strings.ToLower(fn.output.GetExtension()) )
}

// Hash output filename
func (fn *Filenamer) HashIt() {
	f := getMD5Hash(fn.output.GetName())
	fn.output.SetName(f)
}

// Generate filename
func getFormatedTimestamp() string {
	t := time.Now()
	return t.Format("20060102150405")
}

// generate md5 hash from string
func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Get random string with charset with given length
func getRandomStringWithCharset(length int) string {
	r := NewRandomizer()
	return r.Get(length)
}
