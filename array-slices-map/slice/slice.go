package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3} //array tem numero de posicoes definidos
	s1 := []int{1, 2, 3}  //slice nao tem numero de posicoes definidos
	fmt.Println(a1, s1)

	a2 := [5]int{1, 2, 3, 4, 5}

	s2 := a2[1:3] //ele pega a posicao 1 e a 2 nao pega a 3
	fmt.Println(a2, s2)
}
