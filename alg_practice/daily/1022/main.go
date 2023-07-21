package main

import "fmt"

func main() {
	test([]int{1, 1, 1, 5, 5, 5, 7, 20, 7, 7})
}

func findSingleOccuringMap(a []int) int {
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

func findSingleOccuringCount(a []int) int {

	// each index corresponds to a 2's place in binary,
	// eventually this will be filled with columns that are divisible by 3,
	// except for the column that has been added to by the single occuring value
	// which will be the only non divisible by 3 columns. lol
	base2Count := [64]uint8{}

	for _, v := range a {
		for i := 0; i < 64; i++ {
			d := (v >> i) & 1
			base2Count[i] += uint8(d)
		}
	}

	singleOccuringValue := 0
	for i := 0; i < 64; i++ {
		singleOccuringValue += int(base2Count[i]%3) << i
	}

	return singleOccuringValue
}

func test(ints []int) {
	fmt.Printf("input: %v\n", ints)
	r := findSingleOccuringMap(ints)
	r2 := findSingleOccuringCount(ints)
	fmt.Printf("boring result: %d\ncool result: %d\n==========\n", r, r2)
}
