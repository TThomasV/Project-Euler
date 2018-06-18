package main

import (
	"fmt"
)

func main() {
	sumSquares := 0
	squareSum := 0

	for i := 1; i < 101; i++ {
		sumSquares += i * i
	}

	for i := 1; i < 101; i++ {
		squareSum += i
	}
	squareSum *= squareSum

	fmt.Println(squareSum - sumSquares)
}
