package main

import "fmt"

func main() {
	numero(6)
}

func numero(n int) {
	if n > 7 {
		fmt.Println("maior que 7")
	} else if n == 7 {
		fmt.Println("igual que 7")
	} else {
		fmt.Println("menor que 7")
	}
}
