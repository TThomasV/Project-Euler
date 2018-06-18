package main

import (
	"fmt"
	"strconv"
)

func main() {

	max := 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {

			//check palindrome
			result := checkPalindrome(j * i)
			if result > max { // as max always greater than 0 in this case
				max = result // we should be fine
			}
		}
	}

	fmt.Println(max)
}

func checkPalindrome(value int) int {
	if value > 0 { // Just some checking
		value := strconv.Itoa(value) // convert to a string
		runes := []rune(value)       // make a rune slice
		// reverse the string
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		if value == string(runes) { // if the same
			value, _ := strconv.Atoi(value)
			return value
		}

	}

	return -1
}
