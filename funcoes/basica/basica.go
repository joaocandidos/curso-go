package main

import "fmt"

func main() {
	f1()
	f2("joao", "paulo")
	r := f3()
	fmt.Println(r)
	f := f4("um", "dois")
	fmt.Println(f)
	a, b := f5()
	fmt.Println(a, b)
}
func f1() {
	fmt.Println("primeira funcao")
}
func f2(p1 string, p2 string) {
	fmt.Printf("f2: %s %s\n", p1, p2)
}
func f3() string {
	return "f3"
}
func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}
func f5() (string, string) {
	return "retorno 1", "retorno 2"
}
