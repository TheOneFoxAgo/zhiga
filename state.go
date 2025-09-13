package main

import "math/big"

type algorithmState struct {
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
