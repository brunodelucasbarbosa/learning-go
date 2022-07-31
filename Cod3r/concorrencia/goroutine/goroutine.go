package main

import (
	"fmt"
	"time"
)

func say(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// say("João", "Olá", 3)
	// say("Maria", "Tudo bem?", 1)

	go say("João", "Olá", 5)
	go say("Maria", "Tudo bem?", 5)

	time.Sleep(time.Second * 10)
}
