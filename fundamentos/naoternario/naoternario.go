package main

import "fmt"

// nao tem operador ternario na linguagem golang
// podemos fazer assim
func main() {

	fmt.Println(obterResultado(6.2))
}

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "aprovado"
	}
	return "reprovado"
}
