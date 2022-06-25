package main

import (
	"fmt"
)

// dentro da main, criar um for que começa em 0 e vai até 20
// a cada vez que ele rodar, deve imprimir o texto "rodei"
// após a execução das 20 vezes, deve imprimir "finalizado!"

func main() {
	n := 1
	for n <= 20 {
		fmt.Println(n)
		n++
	}
	fmt.Println("finalizado")
}
