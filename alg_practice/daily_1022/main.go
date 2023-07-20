package main

import "fmt"

func main() {
	test([]int{1, 1, 1, 5, 5, 5, 7, 2, 7, 7})
}

func findSingleOccuring(a []int) int {
	m := make(map[int]int)

	for _, v := range a {
		if _, ok := m[v]; !ok {
			m[v] = 0
		}
		m[v]++
	}

	for v, count := range m {
		if count == 1 {
			return v
		}
	}

	return -1
}

func test(ints []int) {
	fmt.Printf("input: %v\n", ints)
	r := findSingleOccuring(ints)
	fmt.Printf("result: %d\n==========\n", r)
}
