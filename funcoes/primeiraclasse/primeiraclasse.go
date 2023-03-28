package main

import "fmt"

func main() {
	sub := func(a1, b1 int) int {
		return a1 - b1
	}
	fmt.Println(soma(25, 41))
	fmt.Println(sub(5, 3))
}

var soma = func(a, b int) int {
	return a + b
}
