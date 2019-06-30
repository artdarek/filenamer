// Copyright 2019 Dariusz Przada. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package Filenamer provides a set of methods for filename manipulation.
// It's a very simple API for adding custom prefixes and suffixes to your
// base filename such as timestamps, random strings etc. Very useful when
// working with file uploads and you need to genereate unique filenames
// for your uploads.
//
// Basic usage example
//
//	source := "filename.jpg"
//	fn := filenamer.New(source)
package filenamer

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"

	"github.com/artdarek/filenamer/filename"
	"github.com/artdarek/filenamer/randomizer"
)

// Filenamer stores input and output file name and it's current separator
type Filenamer struct {
	input     filename.Filename
	output    filename.Filename
	separator string
}

// New creates new Filenamer instance
func New(fileName string) Filenamer {
	input := filename.New(fileName)
	output := filename.New(fileName)
	return Filenamer{input: input, output: output, separator: "_"}
}

// Get returns output filename
func (fn *Filenamer) Get() string {
	return fn.output.Get()
}

// WithTimestamp adds timestamp to filename
func (fn *Filenamer) WithTimestamp() {
	fn.output.SetPrefix(getFormatedTimestamp(), fn.separator)
	fn.output.SetName(fn.output.GetName())
}

// WithRandomPrefix adds unique random string to filename
func (fn *Filenamer) WithRandomPrefix(length int) {
	fn.output.SetPrefix(getRandomStringWithCharset(length), fn.separator)
	fn.output.SetName(fn.output.GetName())
}

// WithRandomSuffix adds unique random string to filename
func (fn *Filenamer) WithRandomSuffix(length int) {
	fn.output.SetSuffix(getRandomStringWithCharset(length), fn.separator)
	fn.output.SetName(fn.output.GetName())
}

// WithCustomPrefix add custom prefix to filename
func (fn *Filenamer) WithCustomPrefix(prefix string) {
	fn.output.SetPrefix(prefix, fn.separator)
	fn.output.SetName(fn.output.GetName())
}

// WithCustomSuffix adds custom suffix to filename
func (fn *Filenamer) WithCustomSuffix(suffix string) {
	fn.output.SetSuffix(suffix, fn.separator)
	fn.output.SetName(fn.output.GetName())
}

// WithExtension adds new extension
func (fn *Filenamer) WithExtension(extension string) {
	fn.output.SetExtension(extension)
}

// WithExtensionRemoved removes extension from file name
func (fn *Filenamer) WithExtensionRemoved() {
	fn.output.SetExtension("")
}

// WithReplacement replaces part of a name with different string
func (fn *Filenamer) WithReplacement(current string, target string) {
	f := strings.Replace(fn.output.GetName(), current, target, -1)
	fn.output.SetName(f)
}

// SeparateWith sets active separator for all feature filename manipulations
func (fn *Filenamer) SeparateWith(separator string) {
	fn.separator = separator
}

// CleanIt cleans all undesired characters from filename
func (fn *Filenamer) CleanIt() {
	fn.output.SetName(strings.Replace(fn.output.GetName(), " ", fn.separator, -1))
	fn.output.SetName(strings.ToLower(fn.output.GetName()))
	fn.output.SetExtension(strings.ToLower(fn.output.GetExtension()))
}

// HashIt hash output filename
func (fn *Filenamer) HashIt() {
	f := getMD5Hash(fn.output.GetName())
	fn.output.SetName(f)
}

// getFormatedTimestamp generates timestamp
func getFormatedTimestamp() string {
	t := time.Now()
	return t.Format("20060102150405")
}

// getMD5Hash generates md5 hash from string
func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// getRandomStringWithCharset gets random string of given length
func getRandomStringWithCharset(length int) string {
	r := randomizer.New()
	return r.Get(length)
}
