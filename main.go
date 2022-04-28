package main

import "fmt"

func main() {
	result := somar(2, 3)
	fmt.Println(result)

}

func somar(numero1 int, numero2 int) int {
	return numero1 + numero2
}
