package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"jose":  1000.00,
		"maria": 2000.00,
		"pedro": 3000.00,
	}
	funcsESalarios["rafael"] = 4500.00

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
