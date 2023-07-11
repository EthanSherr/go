//https://leetcode.com/problems/valid-parentheses/

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/
package main

import "fmt"

type segment struct {
	open  rune
	close rune
	count int
}

func isValid(s string) bool {

	openToClose := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	closeToOpen := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	q := make([]rune, 0)
	for _, r := range s {
		if _, ok := openToClose[r]; ok {
			q = append(q, r)
		} else if open, ok := closeToOpen[r]; ok {
			if len(q) == 0 || open != q[len(q)-1] {
				return false
			}
			q = q[0 : len(q)-1]
		}
	}

	return len(q) == 0
}

func main() {
	arg := "[](({})){}"
	fmt.Println("arg", arg, isValid(arg))
}
