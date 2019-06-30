package randomizer

import (
	"math/rand"
	"time"
)

// Randomizer struct
type Randomizer struct {
	rand    *rand.Rand // Rand pointer
	charset string     // charset to be used
}

// Default charset
const defaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// New create new Randomizer instance
func New() Randomizer {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	return Randomizer{seededRand, defaultCharset}
}

// SetCharset sets charset to use to randomize string
func (rm *Randomizer) SetCharset(charset string) {
	rm.charset = charset
}

// Get returns random string with given length
func (rm *Randomizer) Get(length int) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = rm.charset[rm.rand.Intn(len(rm.charset))]
	}

	return string(b)
}
