package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	"math/bits"
	"os"
	"strings"
)

func parseFString(l *log.Logger, fString string) big.Int {
	f := big.Int{}
	p, _ := f.SetString(fString, 2)
	if p == nil {
		l.Fatalln()
	}
	return f
}
func main() {
	l := log.New(os.Stderr, "zhiga: ", 0)
	var fString, lString string
	flag.StringVar(&fString, "f", "", "Values of function f. Delimeters are allowed. Example: 0101_0010")
	flag.StringVar(&lString, "l", "", "Labels for variables. If provided, the output will be human-readable")
	flag.Parse()
	fString = strings.ReplaceAll(fString, "_", "")
	fLen := uint(len(fString))
	if fLen < 2 {
		l.Fatalln("Number of f values should be at least 2")
	}
	if fLen > 256 {
		l.Fatalln("Number of f values should be at most 256")
	}
	if bits.OnesCount(fLen) != 1 {
		l.Fatalln("Number of f values should be power of 2")
	}
	if lString == "" {
		res := zhigalkin(parseFString(l, fString), fLen)
		fmt.Printf("%0*s\n", fLen, res.Text(2))
	} else {
		labelsAmount := bits.TrailingZeros(fLen)
		if labelsAmount > len(lString) {
			l.Fatalln("Not enough labels")
		}
		for i, c := range lString {
			if i >= labelsAmount {
				break
			}
			fmt.Printf("%c ", c)
		}
		res := zhigalkin(parseFString(l, fString), fLen)
		fmt.Println("| Res:")
		for combination := uint(0); combination < fLen; combination += 1 {
			for mask := fLen >> 1; mask > 0; mask >>= 1 {
				if combination&mask == 0 {
					fmt.Print("0 ")
				} else {
					fmt.Print("1 ")
				}
			}
			fmt.Printf("|  %1b\n", res.Bit(int(combination)))
		}
	}
}
