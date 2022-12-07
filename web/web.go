package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type escritorWeb struct {
}

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	res, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}
	e := escritorWeb{}
	io.Copy(e, res.Body)
}
