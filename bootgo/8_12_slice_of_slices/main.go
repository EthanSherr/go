package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	m := make([][]int, rows)

	for i := 0; i < rows; i++ {
		m[i] = make([]int, cols)
	}

	return m
}

// dont edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
}
