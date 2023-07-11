package main

import (
	"fmt"
)

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 0

	fmt.Printf("debugging actualCostInPennies %f maxCostInPennies %d\n", actualCostInPennies, maxCostInPennies)

	for actualCostInPennies < float64(maxCostInPennies) {
		actualCostInPennies *= costMultiplier
		maxMessagesToSend++
	}
	if int(actualCostInPennies) > maxCostInPennies {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

// don't touch below this line

func test(costMultiplier float64, maxCostInPennies int) {
	maxMessagesToSend := getMaxMessagesToSend(costMultiplier, maxCostInPennies)
	fmt.Printf("Multiplier: %v\n",
		costMultiplier,
	)
	fmt.Printf("Max cost: %v\n",
		maxCostInPennies,
	)
	fmt.Printf("Max messages you can send: %v\n",
		maxMessagesToSend,
	)
	fmt.Println("====================================")
}

func main() {
	test(1.1, 5)
	test(1.3, 10)
	test(1.35, 25)
}
