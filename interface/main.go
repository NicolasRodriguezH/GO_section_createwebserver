package main

import "fmt"

type animal interface {
	mover() string
}

type Perro struct {
}

type Pez struct {
}

type Pajaro struct {
}

func (Perro) mover() string {
	return "Soy un perro y camino"
}

func (Pez) mover() string {
	return "Soy un pez y estoy nadando"
}

func (Pajaro) mover() string {
	return "SOy un pajaro y estoy volando"
}

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {

	p := Perro{}
	moverAnimal(p)
	pe := Pez{}
	moverAnimal(pe)
	pa := Pajaro{}
	moverAnimal(pa)

}
