package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	"math/bits"
	"strings"
	// "slices"
	// "flag"
)

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

func parseFString(fString string) big.Int {
	f := big.Int{}
	p, _ := f.SetString(fString, 2)
	if p == nil {
		log.Fatalln()
	}
	return f
}
func main() {
	var fString, lString string
	flag.StringVar(&fString, "f", "", "Values of function f. Delimeters are allowed. Example: 0101_0010")
	flag.StringVar(&lString, "l", "", "Labels for variables. If provided, the output will be human-readable")
	flag.Parse()
	fString = strings.ReplaceAll(fString, "_", "")
	fLen := uint(len(fString))
	if fLen < 2 {
		log.Fatalln("Number of f values should be at least 2")
	}
	if fLen > 256 {
		log.Fatalln("Number of f values should be at most 256")
	}
	if bits.OnesCount(fLen) != 1 {
		log.Fatalln("Number of f values should be power of 2")
	}
	labelsAmount := bits.LeadingZeros(fLen)
	if lString == "" {
		res := zhigalkin(parseFString(fString), fLen)
		fmt.Println(res.Text(2))
	} else {
		if labelsAmount > len(lString) {
			log.Fatalln("Not enough labels")
		}
		for i, c := range lString {
			if i >= labelsAmount {
				break
			}
			fmt.Print(c, " ")
		}
		fmt.Println("Calculating...")
		res := zhigalkin(parseFString(fString), fLen)
		fmt.Println("| Res:")
		for combination := 0; combination < int(fLen); combination += 1 {
			for mask := 1; mask < int(fLen); mask <<= 1 {
				if combination&mask == 0 {
					fmt.Print("0 ")
				} else {
					fmt.Print("1 ")
				}
			}
			fmt.Printf("|  %b\n", res.Bit(combination))
		}
	}
}

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
