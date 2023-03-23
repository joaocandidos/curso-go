package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	//convertendo int para float
	fmt.Println(x / float64(y))
	//imprime o valor unicod (97 = a)
	fmt.Println(string(97))
	// int para string strconv.Itoa()
	fmt.Println("teste " + strconv.Itoa(12345))
	//de string para int a funcao retorna int e um error
	num, _ := strconv.Atoi("12345")
	fmt.Println(num)
	//converter para bool
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)
}
