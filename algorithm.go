package main

import "math/big"

func zhigalkin(f big.Int, len uint) big.Int {
	initState := algorithmState{f, big.Int{}, big.Int{}}
	initVals := byte(len - 1)
	initCombinations := make([]byte, len)
	for i := range initCombinations {
		initCombinations[i] = byte(i)
	}
	algorithm(initVals, initCombinations, &initState)
	return initState.res
}

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
