package main

import "fmt"

type Vector struct {
	X, Y int
}

var (
	v1 = Vector{1, 2}
	v2 = Vector{Y: 3}
	v3 = Vector{}
	p  = &Vector{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, p, p.X, *p)
	p.X = 555
	fmt.Println(v1, v2, v3, p, p.X, *p)
}
