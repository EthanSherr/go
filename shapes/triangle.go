package main

type triangle struct {
	base   float64
	height float64
}

func (t triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
