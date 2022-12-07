package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instragram.com",
	}

	for _, server := range servidores {
		revisarServidor(server)
	}

	timeLapsed := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion: %s\n", timeLapsed)
}

func revisarServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "no esta dispobible =(")
	} else {
		fmt.Println(servidor, "esta funcionando normalmente =)")
	}
}
