package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Hour()
	switch {
	case t < 12:
		fmt.Println("morning")
	case t < 17:
		fmt.Println("afternoon")
	default:
		fmt.Println("evening")
	}
}
