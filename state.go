package main

import "github.com/bits-and-blooms/bitset"

type algorithmState struct {
	max          uint
	f, done, res bitset.BitSet
}

func (state *algorithmState) isDone(trueVals uint) bool {
	return state.done.Test(uint(trueVals))
}

func (state *algorithmState) setDone(trueVals uint) {
	state.done.Set(uint(trueVals))
}

func (state *algorithmState) Res(trueVals uint) bool {
	return state.res.Test(uint(trueVals))
}

func (state *algorithmState) setRes(trueVals uint, result bool) {
	state.res.SetTo(uint(trueVals), result)
}

func (state *algorithmState) F(trueVals uint) bool {
	return state.f.Test(uint(trueVals))
}
