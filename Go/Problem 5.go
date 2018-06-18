package main

import (
	"fmt"
	"os"
)

func main() {
	testInt := 20 // i cant do maths, this is gonna be really inefficent
	for true {    // bad attempt at a while loop
		for i := 1; i < 22; i++ {
			if testInt%i != 0 { // if remainder not equal 0
				break // breaks the inner loop
			} else if i == 21 { // if we make through to 21, we know it passed all the other ones
				fmt.Println(testInt) // print the value
				os.Exit(0)           // exit the program
			}
		}
		testInt += 20 // has to be divisble by 20, so go in twenties
	}
}
