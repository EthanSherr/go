package main

type rectangle struct {
	side1 float64
	side2 float64
}

func (r rectangle) Area() float64 {
	return r.side1 * r.side2
}
