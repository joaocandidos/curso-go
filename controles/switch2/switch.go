package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("bom dia joao ")
	case t.Hour() < 18:
		fmt.Println("boa tarde joao")
	default:
		fmt.Println("boa noite joao")
	}
}
