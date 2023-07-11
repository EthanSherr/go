// https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/e784c584-f6a7-4c21-bcef-a6c800aa4491/a1eae01c-0a40-47d5-9b98-94fe48199073
package main

import "fmt"

func main() {
	// initialize variables here
	smsSendingLimit := 0
	costPerSMS := 0.0
	hasPermission := false
	username := ""

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}

/*
go's basic types are

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/
