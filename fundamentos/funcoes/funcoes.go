package main

import "fmt"

func main() {
	somar(3, 3)
	mult(2, 5)
}

// funcao com argumentos sem retorno
func somar(a int, b int) {
	fmt.Println(a + b)
}

// funcao com retorno
func mult(a int, b int) int {
	return a * b

}
