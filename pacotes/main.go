package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	r := auxiliar.Somar(2, 45)
	fmt.Println(r)

	erro := checkmail.ValidateFormat("joaocandidos@gmail.com")
	fmt.Println(erro)

}
