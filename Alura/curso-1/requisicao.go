package main

import (
	"fmt"
	"net/http"
)

func main() {

	res, _ := http.Get("http://www.google.com")

	for {
		if res.StatusCode == 200 {
			fmt.Println("O site até o momento está ativo!")
		} else {
			fmt.Println("I")
		}

	}
}
