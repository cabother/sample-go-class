package main

import "fmt"

// DESAFIO 01
// CRIAR UM UM PROBLEMA E UMA SOLUÇÃO, USANDO TUDO QUE JÁ APRENDEU:
// - CRIAR VARIÁVEIS S
// - CRIAR MÉTODOS (FUNC) S
// - RETURN S
// - IF S
// - ELSE S
// - FOR S
// - FMT S
// - UMA OPERAÇÃO MATEMÁTICA S
// - BOOL S
// - ADICIONAR VALIDAÇÃO DE IMPAR OU PAR NO RESULTADO

func main() {
	resultado := -1
	valor := 1

	// chamamos a função "matematica" e guarmos o resultado dela na variavel "resultado"
	resultado = matematica(0, 1)

	verdadeiro := false
	if resultado < 5 {
		verdadeiro = true
	}

	fmt.Println("Resultado Verdadeiro?", verdadeiro)

	for valor <= 3 {
		fmt.Println(resultado)
		valor = valor + 1
	}

	resto := resultado % 2
	if resto == 0 {
		fmt.Println("PAR")
	} else {
		fmt.Println("IMPAR")
	}
}

func matematica(num1 int, num2 int) int {
	if num1 == 3 {
		return 3
	}
	if num1 == 5 {
		return 5
	}
	if num2 == 9 {
		return 9
	}
	if num2 == 4 {
		return 4
	} else {
		return 6
	}
}
