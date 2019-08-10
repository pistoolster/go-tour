package main

import "fmt"

func printSlice(a string, s []int) {
	fmt.Printf("%s len=%d, cap=%d, %v\n", a, len(s), cap(s), s)
}

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)
}
