package main

import "fmt"

func main() {
	i := 1
	var p *int = nil //criando um ponteiro nulo

	/* pegando o endereco de memoria da var i ,
	e atribuindo ao ponteiro nulo */
	p = &i
	fmt.Println(p)
	//utilizado o &i vc pegar o endereco de memoria e nao o valor da variavel
	fmt.Println(&i)

}
