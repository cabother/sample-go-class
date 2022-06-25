package main

import (
	"fmt"
	"strings"
)

func main() {
	nome := "aaaaaabaaaaabaaaaaabaaaabaaa"

	letras := strings.Split(nome, "b")
	for _, letra := range letras {
		fmt.Println(letra)
	}
}
