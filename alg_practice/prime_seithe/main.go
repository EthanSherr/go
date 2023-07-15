package main

import "fmt"

func sieve(nMax int) {

	// if true, it's the mult of something else
	numbers := make([]bool, nMax+1)
	primes := make([]int, 0)

	for n := 2; n <= nMax; n++ {
		if numbers[n] == false {
			primes = append(primes, n)
		}

		times := 2
		mult := times * n

		for mult <= nMax {
			numbers[mult] = true

			times++
			mult *= times * n
		}
	}

	fmt.Printf("primes l.t.e %d are %v", nMax, primes)
}

func main() {
	fmt.Println("hey")
	sieve(10)
}
