package main

import "fmt"

type Vector struct {
	X int
	Y int
}

func main() {
	v := Vector{1, 2}
	fmt.Println(v.X)
	p := &v
	p.Y = 1e9
	fmt.Println(p.Y)
}
