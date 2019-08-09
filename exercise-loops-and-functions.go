package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64) {
	z = x / 2
	for i := 0; ; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, z)
		if z*z-x < math.Pow(10, -15) {
			return
		}
	}
}

func main() {
	a := float64(2)
	fmt.Println("result:", Sqrt(a))
	fmt.Println("math  :", math.Sqrt(a))
}
