package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	"math/bits"
	"strings"
)

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
