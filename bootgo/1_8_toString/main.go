package main

import (
	"fmt"
	"strconv"
)

func main() {
	var username string = "wagslane"
	var password int = 20947382822

	// don't edit below this line
	fmt.Println("Authorization: Basic", username+":"+strconv.Itoa(password))
}
