package main

import (
	"fmt"
	"reflect"
)

func main() {
	//numeros inteiros
	fmt.Println("o tipo é", reflect.TypeOf(500))
	//numeros reais
	fmt.Println("o tipo é", reflect.TypeOf(50.0))
	//boolean
	fmt.Println("o tipo é", reflect.TypeOf(true))
	//string
	fmt.Println("o tipo é", reflect.TypeOf("joao"))
	//tamanho do tipo string
	fmt.Println("o tipo é", len("joaozinho"))

}
