package main

import "fmt"

func adder() func(int) int {
	sum := 0
	time := 0
	return func(x int) int {
		sum += x
		time += 1
		fmt.Println(time)
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
