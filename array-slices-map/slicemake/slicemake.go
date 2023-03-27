package main

import "fmt"

func main() {
	s := make([]int, 10) //funcao make
	//cria um slice apartir de array interno de 10 posicao
	s[9] = 12

	fmt.Println(s)
	//cria um slace de 10 posicoes a partir de um array de 20 posicoes
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))
}
