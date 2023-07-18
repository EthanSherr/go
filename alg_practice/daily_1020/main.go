package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}

func (n node) String() string {
	return fmt.Sprintf("node{value: %d}", n.value)
}

func twoNodesAddUpTo(h *node, sum int) (n1, n2 *node, ok bool) {
	return visitAndLookForComposite(h, sum, make(map[int]*node))
}

func visitAndLookForComposite(n *node, sum int, visited map[int]*node) (n1, n2 *node, ok bool) {
	if n == nil {
		return
	}

	composite := sum - n.value
	if c, ok := visited[composite]; ok {
		return n, c, true
	}
	visited[n.value] = n

	if n1, n2, ok = visitAndLookForComposite(n.left, sum, visited); ok {
		return
	}
	if n1, n2, ok = visitAndLookForComposite(n.right, sum, visited); ok {
		return
	}

	return nil, nil, false
}

func main() {

	hrr := node{value: 15}
	hrl := node{value: 11}
	hr := node{value: 15, left: &hrl, right: &hrr}
	hl := node{value: 5}
	h := node{value: 10, left: &hl, right: &hr}

	sum := 16
	n1, n2, ok := twoNodesAddUpTo(&h, sum)

	fmt.Printf("two nodes add up to thingy? %d : %v %v %v\n", sum, n1, n2, ok)

}
