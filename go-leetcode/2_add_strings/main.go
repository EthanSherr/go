package main

import (
	"fmt"
)

/*
https://leetcode.com/problems/add-strings/
Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

You must solve the problem without using any built-in library for handling large integers (such as BigInteger). You must also not convert the inputs to integers directly.
*/

func main() {
	result := addStrings("99", "999")
	fmt.Println("results =", result)
}

func addStrings(num1 string, num2 string) string {
	index1 := len(num1) - 1
	index2 := len(num2) - 1
	result := make([]byte, 0)

	carry := false
	for index1 > -1 || index2 > -1 {
		sum := 0
		if carry {
			sum = 1
		}
		if index1 > -1 {
			sum += int(num1[index1]) - '0'
		}
		if index2 > -1 {
			sum += int(num2[index2]) - '0'
		}

		result = append(result, byte(sum%10)+'0')
		carry = sum >= 10

		index1--
		index2--
	}

	if carry {
		result = append(result, '1')
	}

	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}

	return string(result)
}
