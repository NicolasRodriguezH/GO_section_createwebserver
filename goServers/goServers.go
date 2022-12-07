package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()

	/* Channel */
	c := make(chan string)

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instragram.com",
	}

	i := 0

	for {
		if i > 2 {
			break
		}
		for _, server := range servidores {
			go revisarServidor(server, c)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-c)
		i++
	}

	timeLapsed := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion: %s\n", timeLapsed)
}

func revisarServidor(servidor string, c chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		c <- servidor + " no se encuentra dispobible"
	} else {
		c <- servidor + " esta funcionando"
	}
}
