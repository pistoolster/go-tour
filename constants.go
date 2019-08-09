package main

import "fmt"

const Pi float64 = 3.14

func main() {
	const World string = "world"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	//World = ''
	const Truth bool = true
	fmt.Println("Go rules?", Truth)
}
