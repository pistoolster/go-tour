package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a)
	fmt.Println(a[0], a[1])

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	c := primes[2:5]
	fmt.Println(c)

	c[1] = 666

	fmt.Println(primes, s, c)
}
