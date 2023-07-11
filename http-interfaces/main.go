package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println("hello ", len(bs), "\n\n", string(bs))
	return len(bs), nil
}

func main() {
	resp, httpErr := http.Get("http://google.com")
	fmt.Println("resp is", resp)
	if httpErr != nil {
		fmt.Println("Error:", httpErr)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}
