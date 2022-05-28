package main

import "fmt"

func main() {
	//Declarar uma variável int com valor inicial 15
	valorInteiro := 15
	fmt.Println(valorInteiro)

	//Declarar uma variável string com valor inicial "Luis Otávio"
	nome := "LUIS OTAVIO"
	fmt.Println(nome)

	//Declarar uma variável bool com valor inicial true
	ball := true
	fmt.Println(ball)

	//Declarar valiável do tipo float (com valor inicial 22.77)
	valorFloat := 22.77
	fmt.Println(valorFloat)

	valorSomado := valorFloat + float64(valorInteiro)
	fmt.Println(valorSomado)

	fmt.Println("Hello World - Variáveis")
}
