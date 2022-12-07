package main

import "fmt"

func main() {
	m1 := make(map[int]string)

	m1[1] = "Hola"
	fmt.Println(m1)
	fmt.Println(m1[1])
}
