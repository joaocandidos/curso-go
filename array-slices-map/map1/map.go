package main

import "fmt"

func main() {
	aprovados := make(map[int]string) //chave: valor

	aprovados[12345678] = "maria"
	aprovados[11345678] = "joao"
	aprovados[11145678] = "paulo"
	fmt.Println(aprovados)

}
