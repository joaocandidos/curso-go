package main

import "fmt"

func main() {
	fmt.Println(tipo(5))
	fmt.Println(tipo(4.9))
	fmt.Println(tipo("joao"))
	fmt.Println(tipo(func() {}))

}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "nao sei o tipo"
	}
}
