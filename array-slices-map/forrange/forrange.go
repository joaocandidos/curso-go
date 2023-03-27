package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // [...] copilador conta quantos elementos vai ter no array

	for i, numero := range numeros {
		fmt.Printf("indice [%d] =  %d\n", i, numero)
	}

	//pegar so os valores sem mostrar o indice
	for _, num := range numeros {
		fmt.Print(num, "  ")
	}
}
