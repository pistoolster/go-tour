package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, java, python)
	var x, y int
	fmt.Println(x == y, &x == &y, &x, &y, &x == nil)
}
