package main

import "fmt"

func fibonacci() func() int { // 1,1,2,3,5,8....
	a, b := 0, 1
	return func() (f int) {
		f = a
		a, b = b, a+b
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
