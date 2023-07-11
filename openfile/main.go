package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./cool.txt")

	if err != nil {
		log.Fatal(err)
	}
	//do something

	data := make([]byte, 5)
	for {
		data = data[:cap(data)]
		n, err := f.Read(data)
		if err == io.EOF {
			break
		}
		data = data[:n]
	}

	fmt.Printf("bytes: %s\n", string(data))

	// r4 := bufio.NewReader(f)
	// b4, err := r4.Peek(5)
	// fmt.Printf("5 bytes: %s\n", string(b4))
	// check(err)
	// f.Close()

}
