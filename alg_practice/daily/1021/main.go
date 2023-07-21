package main

import (
	"fmt"
	"strconv"
)

func main() {
	inputs := []string{"11101010"}
	for _, v := range inputs {
		i, err := strconv.ParseUint(v, 2, 8)
		if err != nil {
			fmt.Println(err)
			continue
		}
		test(uint8(i))
	}
}

func flipBits(n uint8) uint8 {
	var r uint8 = 0

	for i := 0; i < 8; i++ {
		b := (n >> i) & 0x1
		if i%2 == 0 {
			r += b << (i + 1)
		} else {
			r += b << (i - 1)
		}
		fmt.Printf("index %d has value %d\n shifted: %b\n", i, b, n>>i)
	}

	return r
}

func flipBitsBetter(n uint8) uint8 {
	odsMask, _ := strconv.ParseUint("10101010", 2, 8)
	evsMask := odsMask >> 1 // 01010101

	return ((n & uint8(odsMask)) >> 1) + ((n & uint8(evsMask)) >> 1)
}

func test(n uint8) {

	fmt.Printf("Swap even and odd bits of %d\n", n)

	r := flipBits(n)
	rb := flipBitsBetter(n)
	fmt.Printf("bad %b => %b\n", n, r)
	fmt.Printf("gd %b => %b\n", n, rb)

	fmt.Println("=================")
}

// 11101010
// 11010101
//
