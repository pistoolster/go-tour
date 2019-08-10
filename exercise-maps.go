package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WorldCount(s string) map[string]int {
	m := make(map[string]int)
	fmt.Println("Fields--->", strings.Fields(s))
	for _, v := range strings.Fields(s) {
		fmt.Println(v)
		m[v] += 1
	}
	return m
}

func main() {
	wc.Test(WorldCount)
}
