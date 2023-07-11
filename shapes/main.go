package main

import "fmt"

func main() {
	r := rectangle{side1: 25.0, side2: 30.0}
	printShape(r)

	c := circle{radius: 2.0}
	printShape(c)
}

func printShape(s Shape) {
	fmt.Printf("%T: %v\n", s, s.Area())
}
