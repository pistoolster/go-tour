package main

import (
	"fmt"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Scale(f float64) {
	v.x *= f
	v.y *= f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v)
}