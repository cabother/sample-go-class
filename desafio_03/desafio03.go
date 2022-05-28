package main

import "fmt"

// Criar função chamada "somarValores" com retorno do tipo float
// Função acima deve receber tres valores float
// Efetuar a soma dos 3 valores e guardar em uma variável chamada "resultadoDaSoma"
// criar if vendo se o valor da soma é maior que 1000, se sim, printar "SIM", se nao, printar "NAO"
// retornar a variável resultadoSoma

// chamar a função somarValores dentro da funcao main
// e fazer um if vendo se o valor de resultado é menor que 100, printar "menor" ou "maior"
// printar o resultado

func main() {
	calcular := somarValores(90.03, 150.33, 1000.33)
	if calcular < 100 {
		fmt.Println("menor")
	} else {
		fmt.Println("maior")
	}

	resultado := fmt.Sprintf("%f", calcular)
	fmt.Println(resultado)
}

func somarValores(num1 float64, num2 float64, num3 float64) float64 {
	resultadoDaSoma := num1 + num2 + num3
	if resultadoDaSoma > 1000 {
		fmt.Println("SIM")
	} else {
		fmt.Println("NÃO")
	}
	return resultadoDaSoma
}
