package main

import "fmt"

func main() {
	x := 25
	fmt.Println(x)
	fmt.Println(&x)
	y := &x
	cambiarValor(*y)
	fmt.Println(y)
	fmt.Println(*y)
}

func cambiarValor(a int) {
	fmt.Println(&a)
	a = 34
}
