package main

import "fmt"

func main() {
	fmt.Println("Hello World - Funções Matemáticas")
	resultadosomar := somar(2, 3)
	fmt.Println("soma", resultadosomar)
	resultadosubitrair := subtrair(3, 10)
	fmt.Println("sub", resultadosubitrair)
	resultadomultiplicar := multiplicar(5, 5)
	fmt.Println("multiplicar", resultadomultiplicar)
	resultadocalcular := calcular(2, 1, "menos")
	fmt.Println("calcular", resultadocalcular)
}

// Criar função somar
// - Deve receber dois números (numero1 e numero2)
// devolver (retornar) a soma dos dois
func somar(numero1 int, numero2 int) int {
	return numero1 + numero2
}

// Criar função subtrair
// - Deve receber dois números (numero1 e numero2)
// e devolver (retornar) a subtração do primeiro menos o segundo
func subtrair(numero1 int, numero2 int) int {
	return numero1 - numero2
}

// Criar função multiplicar
// - Deve receber dois números (numero1 e numero2)
// e devolver (retornar) a subtração do primeiro menos o segundo
func multiplicar(numero1 int, numero2 int) int {
	return numero1 * numero2
}

// Criar função chamada calcular
// Deverá receber dois numeros (numero1, numero2) e também uma string chamada (tipo)
// SE o tipo for igual a "mais", deve devolver a soma dos dois números
// SE o tipo for igual a "menos" deve devolver a subtração dos dois numeros
// SE NAO for igual aos valores acima, devolver o segundo numero recebido
func calcular(numero1 int, numero2 int, tipo string) int {
	if tipo == "mais" {
		return numero1 + numero2
	}
	if tipo == "menos" {
		return numero1 - numero2
	} else {
		return numero2
	}

}
