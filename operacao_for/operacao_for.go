package main

import (
	"fmt"
	"strconv"
)

// O QUE APRENDER?
// Aprender o uso de FOR
//
// for(i :=0; i< 10; i++)
//
// contador := 100
// for contador >= 0 {
// 	fmt.Println("A quantidade é: ", contador)
// 	contador = contador - 1
// }

//DESAFIO!
// Seu for deve percorrer 100 x
// Declarar uma variavel chamada contador, que deve começar em 100 e terminar em 0
// A cada vez que percorrer ele deve subtrair 1 do contador

// Tabuada do 9
// 9 x 1 = 9
// 9 x 2 = 18
// 9 x 3 = 27
// ...

// numero := 1
// resultado := 0
// for numero <= 10 {
// 	resultado = 9 * numero
// 	fmt.Println("9 x " + strconv.Itoa(numero) + " = " + strconv.Itoa(resultado))
// 	numero++
// }

func main() {
	numero := 10
	resultado := 0
	for numero >= -10 {
		resultado = 10 * numero
		fmt.Println("10 x " + strconv.Itoa(numero) + " = " + strconv.Itoa(resultado))
		numero--
	}
}
