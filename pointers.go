package main

import "fmt"

func main() {
	i, j := 123, 456
	p := &i
	fmt.Println(*p)
	fmt.Println(*&i)
	*p = 321
	fmt.Println(i)

	p = &j
	*p = *p / 2
	fmt.Println(j)
}
