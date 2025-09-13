package main

import (
	"fmt"
	"math/big"
	// "slices"
	// "flag"
)

func algorithm(trueVals byte, combinations2check []byte, state *algorithmState) {
	// sink
	for mask := byte(1); mask != 0; mask <<= 1 {
		next := trueVals ^ mask
		if next > trueVals {
			continue
		} else if !state.isDone(next) {
			nextCombinations := make([]byte, len(combinations2check)/2)
			idx := 0
			for _, c := range combinations2check {
				if c&mask == 0 {
					nextCombinations[idx] = c
					idx += 1
				}
			}
			algorithm(next, nextCombinations, state)
		}
	}
	// float
	res := state.F(trueVals)
	for _, c := range combinations2check {
		if state.Res(c) {
			res = !res
		}
	}
	state.setDone(trueVals)
	state.setRes(trueVals, res)
}

func main() {

}

type algorithmState struct {
	len          int
	f, done, res big.Int
}

func (state *algorithmState) isDone(trueVals byte) bool {
	return state.done.Bit(int(trueVals)) == 1
}

func (state *algorithmState) setDone(trueVals byte) {
	state.done.SetBit(&state.done, int(trueVals), 1)
}

func (state *algorithmState) Res(trueVals byte) bool {
	return state.res.Bit(int(trueVals)) == 1
}

func (state *algorithmState) setRes(trueVals byte, result bool) {
	var intResult uint = 0
	if result {
		intResult = 1
	}
	state.done.SetBit(&state.done, int(trueVals), intResult)
}

func (state *algorithmState) F(trueVals byte) bool {
	return state.f.Bit(int(trueVals)) == 1
}
