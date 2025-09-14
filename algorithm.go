package main

import "github.com/bits-and-blooms/bitset"

func zhigalkin(f bitset.BitSet, len uint) bitset.BitSet {
	initVals := len - 1
	initState := algorithmState{initVals, f, bitset.BitSet{}, bitset.BitSet{}}
	algorithm(initVals, &initState)
	return initState.res
}

func algorithm(trueVals uint, state *algorithmState) {
	// sink
	var minor, mask, next uint
	for rem := trueVals; rem != 0; rem &= minor {
		minor = rem - 1
		mask = rem & ^minor
		next = trueVals ^ mask
		if !state.isDone(next) {
			algorithm(next, state)
		}
	}
	// float
	res := state.F(trueVals)
	banned := trueVals ^ state.max
	incr := banned + 1
	idx := uint(0)
	for {
		if state.Res(idx) {
			res = !res
		}
		idx = trueVals & (idx + incr)
		if idx == 0 {
			break
		}
	}
	state.setDone(trueVals)
	state.setRes(trueVals, res)
}
