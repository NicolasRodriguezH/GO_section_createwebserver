package main

import "fmt"

type Perro struct {
}

type Pez struct {
}

type Pajaro struct {
}

func (Perro) Caminar() string {
	return "Soy un perro y camino"
}

func (Pez) Nadar() string {
	return "Soy un pez y estoy nadando"
}

func (Pajaro) Vuela() string {
	return "SOy un pajaro y estoy volando"
}

func moverPerro(p Perro) {
	fmt.Println(p.Caminar())
}

func moverPez(p Pez) {
	fmt.Println(p.Nadar())
}

func moverPajaro(p Pajaro) {
	fmt.Println(p.Vuela())
}

func main() {

	p := Perro{}
	moverPerro(p)
	pe := Pez{}
	moverPez(pe)
	pa := Pajaro{}
	moverPajaro(pa)

}
