package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello")

	test("12:45")
	test("12:59")
	test("1:00")
}

func angleBetweenHands(hhmm string) (int, error) {

	splitHhMm := strings.Split(hhmm, ":")
	if len(splitHhMm) != 2 {
		return 0, errors.New(fmt.Sprintf("hhmm \"%s\" is invalid format", hhmm))
	}

	hh, err1 := strconv.Atoi(splitHhMm[0])
	if err1 != nil {
		return 0, err1
	}
	mm, err2 := strconv.Atoi(splitHhMm[1])
	if err2 != nil {
		return 0, err2
	}

	if !(hh > 0 && hh < 13) {
		return 0, errors.New(fmt.Sprintf("hh \"%d\" is out of range", hh))
	}
	if !(mm >= 0 && mm < 60) {
		return 0, errors.New(fmt.Sprintf("mm \"%d\" is out of range", mm))
	}

	pMM := float64(mm) / 60.0
	pHH := (float64(hh) + pMM) / 12.0

	return int((pHH - pMM) * 360), nil
}

func test(hhmm string) {
	fmt.Printf("hhmms: %s\n", hhmm)
	r, err := angleBetweenHands(hhmm)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %d\n", r)
	fmt.Println("=====")
}
