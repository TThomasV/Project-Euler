package main

import (
	"fmt"
)

func main() {
	// using go converted version of https://stackoverflow.com/a/412942
	var factors []int
	d := 2
	n := 600851475143
	for n > 1 {
		for n%d == 0 {
			factors = append(factors, d)
			n /= d
		}

		d = d + 1

		if d*d > n {
			if n > 1 {
				factors = append(factors, n)
				break
			}
		}
	}

	// Modified version of https://play.golang.org/p/vhHmjhOMEo
	// Uses biggest only

	biggest := factors[0]

	for _, v := range factors {
		if v > biggest {
			biggest = v
		}
	}

	fmt.Println(biggest)
}
