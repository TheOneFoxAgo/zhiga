package main

import (
	"testing"

	"github.com/bits-and-blooms/bitset"
)

func TestBasicF(t *testing.T) {
	baseTest(t, 4, 0b0101, 0b0011)
}

func TestAdvancedF(t *testing.T) {
	baseTest(t, 16, 0b0001_0000_1100_0100, 0b0111_1100_1000_1100)
}

func baseTest(t *testing.T, len uint, in uint64, out uint64) {
	var f, want bitset.BitSet
	f.SetBitsetFrom([]uint64{in})
	want.SetBitsetFrom([]uint64{out})
	res := zhigalkin(f, len)
	if res.Equal(&want) {
		t.Errorf("Want: %*s, got: %*s", len, want.DumpAsBits(), len, res.DumpAsBits())
	}
}
