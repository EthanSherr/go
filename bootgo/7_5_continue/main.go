package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n == 2 {
		return true
	}

	for i := 3; i < n; i++ {
		if n%i == 0 {
			fmt.Println("found!")
			return false
		}
	}

	return true
}

func printPrimes(max int) {
	for n := 2; n < max+1; n++ {
		if isPrime(n) {
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
	// test(20)
	// test(30)
}
