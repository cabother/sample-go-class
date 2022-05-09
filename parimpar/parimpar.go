package main

import "fmt"

func main() {
	resultado := 6
	resto := resultado % 2
	parOuImpar := ""

	if resto == 0 {
		parOuImpar = "par"
	} else {
		parOuImpar = "impar"
	}

	fmt.Println("TIPO:", parOuImpar)
	fmt.Println("RESTO:", resto)
}
