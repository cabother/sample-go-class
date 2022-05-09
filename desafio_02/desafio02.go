package main

import (
	"fmt"
)

func main() {
	pate := 3
	resto := pate % 2
	if resto == 0 {

		fmt.Println("o resutado é par")
	} else {
		fmt.Println("o resutado é impar")
	}

	fmt.Println(resto)
}
