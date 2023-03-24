package main

import "fmt"

func main() {
	b := 2

	b += 2
	fmt.Println("2 + 2 ", b)
	b = 2
	b -= 2
	fmt.Println("2 - 2 ", b)
	b = 2
	b /= 2
	fmt.Println("2 / 2 ", b)
	b = 2
	b *= 2
	fmt.Println("2 * 2 ", b)
}
