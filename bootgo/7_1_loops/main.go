package main

import (
	"fmt"
)

func bulkSend(numMessages int) float64 {
	// ? using a for loop altho mult is better
	sum := 1.0
	for i := 0; i < numMessages-1; i++ {
		sum += 0.01
	}

	return sum
}

// don't edit below this line

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func main() {
	test(1)
	test(2)
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)
}
