package main

import "fmt"

func main() {

	index := 0
	for index < 10 {
		fmt.Println("index", index)
		addOne(&index)
	}
}

func addOne(index *int) {
	*index = *index + 1
}

// func getNext(a []int, index *int) (int, bool) {
// 	if *index >= len(a) {
// 		return 0, false
// 	}
// 	value := a[*index]
// 	index = *index + 1

// 	return value, true
// }
