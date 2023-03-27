package main

import "fmt"

func main() {
	var notas [3]float64
	notas[0], notas[1], notas[2] = 5.5, 4.4, 3.3

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]

		fmt.Println(total)
	}
	media := total / float64(len(notas))
	fmt.Printf("media = %.2f\n", media)
}
