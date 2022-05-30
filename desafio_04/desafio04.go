package main

import "fmt"

// Criar função chamada "fazerJanela", sua função deve receber a altura e largura e
// devolver a multiplicação da altura por largura

func main() {
	novajanela := fazerJanela(10.99900, 100.22)
	fmt.Println(novajanela)

}

func fazerJanela(alt float32, largu float32) float32 {
	return alt * largu
}
