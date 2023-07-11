package main

import "fmt"

func fizzbuzz() {
	for i := 0; i < 100; i++ {
		mult3 := i%3 == 0
		mult5 := i%5 == 0

		if mult3 && mult5 {
			fmt.Println("fizzbuzz")
		} else if mult3 {
			fmt.Println("fizz")
		} else if mult5 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz()
}
