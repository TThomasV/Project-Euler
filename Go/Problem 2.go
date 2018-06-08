package main

import (
	"fmt"
)

func main() {
	// some really bad fibonacci code
	// martin would be crying right now
	fib1 := 1
	fib2 := 2
	temp := 0
	sum := 0
	for (fib2) < 4000000 {
		if (fib2)%2 == 0 {
			sum += fib2
		}
		temp = fib1
		fib1 = fib2
		fib2 = fib1 + temp

	}
	fmt.Println(sum)

}
