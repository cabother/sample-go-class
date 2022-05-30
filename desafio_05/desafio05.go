package main

import (
	"fmt"
)

// fazerfunc que se cha mae janela azul
// recbe a cor e o tamanho
// se for verde eu que ro multiplique por1000
// senao tirar 100 do tamanho
// deve retor a ocor e a largura

func main() {
	nome := pegar()
	fmt.Println(nome)

	nomeCor, result, newbool := janelaAzul("verde", 32.99, 1000)

	fmt.Println(nomeCor)
	fmt.Println(result)
	fmt.Println(newbool)
}

func janelaAzul(cor string, largura float32, altura float32) (string, float32, int64) {
	var resultado float32

	if cor == "verde" {
		resultado = largura * 1000
	} else {
		resultado = largura - 100
	}

	return cor, resultado, 1000
}

func pegar() string {
	return "otavio"
}
