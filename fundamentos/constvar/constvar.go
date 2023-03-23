package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	//forma reduzida de criar uma var

	area := PI * m.Pow(raio, 2)
	fmt.Println("A area da circunferencia é :", area)

	//declaracao de variaveis em bloco
	var (
		a = 1
		b = 3
	)
	//declaracao de constantes em bloco
	const (
		c = 11
		d = 22
	)
	//imprimindo as constantes e as variaveis
	fmt.Println(a, b, c, d)
	//declaracao reduzida de variaveis

	q, w, e := 3, false, " verdadeiro "

	fmt.Print(q, w, e)

}
