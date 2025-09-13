package main

import (
	"math/big"
	"testing"
)

func TestBasicF(t *testing.T) {
	baseTest(t, 4, 0b0101, 0b0011)
}

func TestAdvancedF(t *testing.T) {
	baseTest(t, 16, 0b0001_0000_1100_0100, 0b0111_1100_1000_1100)
}

func baseTest(t *testing.T, len uint, in int64, out int64) {
	f := big.NewInt(in)
	res := zhigalkin(*f, len)
	if res.Cmp(big.NewInt(out)) != 0 {
		t.Errorf("Want: %b, got: %b", out, res.Int64())
	}
}
