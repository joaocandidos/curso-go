package main

import "runtime/debug"

func main() {
	f1()
}

func f3() {
	debug.PrintStack() //inprime a pilha de execucao
}
func f2() {
	f3()
}
func f1() {
	f2()
}
