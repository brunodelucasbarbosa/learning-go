package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("!=:", 3 != 2)
	fmt.Println("<=:", 3 <= 2)
	fmt.Println(">=:", 3 >= 2)
	fmt.Println("<:", 3 < 2)
	fmt.Println(">:", 3 > 2)

	data := time.Unix(0, 0)
	data2 := time.Unix(0, 0)

	fmt.Println(data)
	fmt.Println(data2)

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}

	fmt.Println(p1 == p2)

}
