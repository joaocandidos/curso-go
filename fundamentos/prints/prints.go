package main

import "fmt"

func main() {

	fmt.Println("pula")
	fmt.Print("linha")

	x := 9.1
	a := 1200
	s := 11
	//sprint transforma em string
	xs := fmt.Sprint(x)

	fmt.Println("o valor de x é:" + xs)
	//print formatado controlasndo o numero de casas decimais
	fmt.Printf("o valor de x é: %.3f", x)
	fmt.Printf("\n%v %v %v", x, a, s)

}
