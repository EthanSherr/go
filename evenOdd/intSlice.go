package main

// where n is exclusive
func newIntSlice(n int) []int {
	intSlice := []int{}
	for i := 0; i < n; i++ {
		intSlice = append(intSlice, i)
	}
	return intSlice
}
