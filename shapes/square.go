package main

type square struct {
	sideLength float64
}

func (s square) Area() float64 {
	return s.sideLength * s.sideLength
}
