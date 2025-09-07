package main

import (
	"fmt"
)

// the goal of the exercise is using break/continue thats why
// no need to puke, when looking at the code
func printPrimes(max int) {
	for n := 2; n <= max; n++ {
		isPrime := true
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
				// or continue outer, but we didn't do labels yet
			}
		}
		if isPrime {
			fmt.Println(n)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
