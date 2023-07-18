package main

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

func (p point) String() string {
	return fmt.Sprintf("[%d, %d]", p.x, p.y)
}

type rect struct {
	p1 point
	p2 point
}

func (r rect) String() string {
	return fmt.Sprintf("[%s, %s]", r.p1.String(), r.p2.String())
}

func (r rect) points() [4]point {
	p := [4]point{r.p1, {r.p1.x, r.p2.y}, r.p2, {r.p2.x, r.p1.y}}
	return p
}

func (r rect) contains(p point) bool {
	xInside := !(r.p1.x > p.x || r.p2.x < p.x)
	yInside := !(r.p1.y > p.y || r.p2.y < p.y)
	return xInside && yInside
}

func (r rect) area() int {
	return (r.p2.x - r.p1.x) * (r.p2.y - r.p1.y)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (r1 rect) overlapAreaBadder(r2 rect) int {

	myMap := make(map[string]bool)

	for x1 := r1.p1.x; x1 < r1.p2.x; x1++ {
		for y1 := r1.p1.y; y1 < r1.p2.y; y1++ {
			myMap[fmt.Sprintf("%d,%d", x1, y1)] = true
		}
	}

	area := 0
	for x1 := r2.p1.x; x1 < r2.p2.x; x1++ {
		for y1 := r2.p1.y; y1 < r2.p2.y; y1++ {
			if myMap[fmt.Sprintf("%d,%d", x1, y1)] {
				area++
			}
		}
	}

	return area
}

func (r1 rect) overlapAreaBetter(r2 rect) int {
	intercepts := make([]point, 0)

	for _, p := range r1.points() {
		if r2.contains(p) {
			intercepts = append(intercepts, p)
		}
	}

	for _, p := range r2.points() {
		if r1.contains(p) {
			intercepts = append(intercepts, p)
		}
	}

	fmt.Printf("intercepts %v\n", intercepts)

	p1 := point{math.MaxInt, math.MaxInt}
	p2 := point{math.MinInt, math.MinInt}

	for _, p := range intercepts {
		p1.x = min(p1.x, p.x)
		p1.y = min(p1.y, p.y)

		p2.x = max(p2.x, p.x)
		p2.y = max(p2.y, p.y)
	}

	interceptRect := rect{p1: p1, p2: p2}
	fmt.Printf("interceptRect %v\n", interceptRect)

	return interceptRect.area()
}

func main() {
	test(
		rect{p1: point{1, 2}, p2: point{3, 4}},
		rect{p1: point{1, 2}, p2: point{3, 4}},
	)
}

func test(r1, r2 rect) {
	fmt.Printf("debug test \nr1: %v \nr2: %v\n", r1, r2)
	basicResult := r1.overlapAreaBadder(r2)
	goodResult := r1.overlapAreaBetter(r2)
	fmt.Printf("basicResult: %d, goodResult: %d\n", basicResult, goodResult)
}
