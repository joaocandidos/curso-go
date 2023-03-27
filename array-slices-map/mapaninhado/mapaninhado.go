package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{

		"G": {
			"Gabriela Silva": 12345678.00,
			"Gabriel gomes":  12345678.99,
		},
		"J": {
			"Jabriela Silva": 12345678.00,
			"Jabriel gomes":  12345678.99,
		},
	}
	//delete(funcsPorLetra, "G")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
