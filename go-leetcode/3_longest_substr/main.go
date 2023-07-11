// https://leetcode.com/problems/longest-substring-without-repeating-characters/

package main

import "fmt"

func main() {
	res := lengthOfLongestSubstring("abcabcbb")
	fmt.Println("res", res)
}

func lengthOfLongestSubstringOld(s string) int {
	maxLength := 0
	start := 0
	for start < len(s) {
		runeMap := make(map[byte]bool)
		tail := start
		for tail < len(s) && !runeMap[s[tail]] {
			runeMap[s[tail]] = true
			tail++
		}
		sequenceLength := tail - start
		if maxLength < sequenceLength {
			maxLength = sequenceLength
		}
		start++
	}

	return maxLength
}
func lengthOfLongestSubstring(s string) int {
	maxLength := 0

	sequenceRecord := make(map[byte]bool)

	start := 0
	end := 0
	for end < len(s) {
		b := s[end]
		if sequenceRecord[b] {
			// advance start, removing sequenceRecord as I go, until I get to a b
			for s[start] != b {
				delete(sequenceRecord, s[start])
				start++
			}
			start++
		}
		sequenceRecord[b] = true
		length := end - start + 1
		if maxLength < length {
			maxLength = length
		}

		end++
	}

	return maxLength
}
