package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("42", myAtoi("42"))
	fmt.Println(math.MaxInt32, math.MinInt32)
	fmt.Println(5<<1, math.MinInt32)

	fmt.Println("9223372036854775808", myAtoi("9223372036854775808"))
	// fmt.Println("     - 42", myAtoi("     - 42"))
}

func myAtoi(s string) int {
	num := 0
	index := 0
	// trim whitespace
	for index < len(s) && s[index] == ' ' {
		index++
	}
	if index == len(s) {
		return num
	}

	sign := 1
	if s[index] == '+' {
		index++
	} else if s[index] == '-' {
		sign = -1
		index++
	}

	for index < len(s) && s[index] >= '0' && s[index] <= '9' {
		digit := (int)(s[index] - '0')

		num = num*10 + digit
		index++
	}
	result := sign * num
	if result < math.MinInt32 {
		return math.MinInt32
	}
	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	return sign * num
}
