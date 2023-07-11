package main

import "fmt"

func main() {
	for _, index := range newIntSlice(11) {
		isOdd := index%2 == 1
		msg := "even"
		if isOdd {
			msg = "odd"
		}
		fmt.Printf("%v %v\n", index, msg)
	}
}
