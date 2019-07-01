package randomizer

import (
	"testing"
)

func TestRandomizer_Get(t *testing.T) {
	rnd := New()
	if len(rnd.Get(10)) != 10 {
		t.Error("Expected length to be 10")
	}
}

func TestRandomizer_SetCharset(t *testing.T) {
	rnd := New()
	rnd.SetCharset("ab")
	if len(rnd.Get(1)) != 1 && ( rnd.Get(1) != "a" || rnd.Get(1) != "b" ) {
		t.Error("Expected length to be 1 and equals to a or b")
	}
}