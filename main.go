package main

import (
	"flag"
	"fmt"
	"github.com/bits-and-blooms/bitset"
	"log"
	"math/bits"
	"os"
	"strings"
)

func parseFString(l *log.Logger, fString string, reverse bool) (result bitset.BitSet) {
	for i, c := range fString {
		switch c {
		case '1':
			if reverse {
				result.Set(uint(i))
			} else {
				result.Set(uint(len(fString) - i - 1))
			}
		case '0':
			continue
		default:
			l.Fatalln("Incorrect string format")
		}
	}
	return
}
func main() {
	l := log.New(os.Stderr, "zhiga: ", 0)
	var tFlag, fFlag, lFlag string
	flag.StringVar(&fFlag, "f", "", "Values of function f. Delimeters are allowed. Example: 0101_0010")
	flag.StringVar(&tFlag, "t", "", "Values of function f but in reverse order. Delimeters are allowed. Example: 0100_1010")
	flag.StringVar(&lFlag, "l", "", "Labels for variables. If provided, the output will be human-readable")
	flag.Parse()

	reverse := fFlag == ""
	var fString string
	if reverse {
		fString = strings.ReplaceAll(tFlag, "_", "")
	} else {
		fString = strings.ReplaceAll(fFlag, "_", "")
	}
	fLen := uint(len(fString))
	if fLen < 2 {
		l.Fatalln("Number of f values should be at least 2")
	}
	if fLen > (1 << 16) {
		l.Fatalf("Number of f values should be at most %d\n", 1<<16)
	}
	if bits.OnesCount(fLen) != 1 {
		l.Fatalln("Number of f values should be power of 2")
	}
	if lFlag == "" {
		res := zhigalkin(parseFString(l, fString, reverse), fLen)
		fmt.Printf("%0*s\n", fLen, res.DumpAsBits())
	} else {
		labelsAmount := bits.TrailingZeros(fLen)
		if labelsAmount > len(lFlag) {
			l.Fatalln("Not enough labels")
		}
		for i, c := range lFlag {
			if i >= labelsAmount {
				break
			}
			fmt.Printf("%c ", c)
		}
		res := zhigalkin(parseFString(l, fString, reverse), fLen)
		fmt.Println("| Res:")
		for combination := uint(0); combination < fLen; combination += 1 {
			for mask := fLen >> 1; mask > 0; mask >>= 1 {
				if combination&mask == 0 {
					fmt.Print("0 ")
				} else {
					fmt.Print("1 ")
				}
			}
			if res.Test(combination) {
				fmt.Println("|  1")
			} else {
				fmt.Println("|  0")
			}
		}
	}
}
