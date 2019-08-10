package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{1, 2, 3, 4, 6, 7}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[2:4]
	printSlice(s)

	//s = s[:7]
	//printSlice(s)

	var t []int
	printSlice(t)
	if t == nil {
		fmt.Println("nil")
	}
}
