package filenamer

import (
	"math/rand"
	"time"
)

type Randomizer struct {
	rand *rand.Rand
	charset string
}

const defaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewRandomizer() Randomizer {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	return Randomizer{seededRand, defaultCharset}
}

func (rm *Randomizer) SetCharset(charset string) {
	rm.charset = charset
}

// Get random string with given length
func (rm *Randomizer) Get(length int) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = rm.charset[rm.rand.Intn(len(rm.charset))]
	}

	return string(b)
}