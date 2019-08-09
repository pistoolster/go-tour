package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	a := 123
	a = a + 1
	return
}

func main() {
	fmt.Println(split(35))
}
