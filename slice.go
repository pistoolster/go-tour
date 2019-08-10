package main

import "fmt"

func main() {
	q := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(q)

	s := []struct {
		x string
		y int
	}{
		{"hello", 1},
		{"world", 2},
		{x: "test"},
	}

	fmt.Println(s)

	fmt.Println(q[:], q[2:4], q[:6])
}
