package main

import (
	"fmt"
	"reflect"
)

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "yes!"
	fmt.Println(c, python, java, i, j)
	fmt.Println(reflect.TypeOf(c))
}
