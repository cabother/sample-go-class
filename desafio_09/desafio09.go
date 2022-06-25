package main

import (
	"fmt"
	"time"
)

// chamar a funcção validar passando um nome
// usando switch validar se o nome é igual a otavio, se sim, retornar true
// se não, retornar false
// e default retornar false

func main() {
	retorno := validar("otavio")
	fmt.Println(retorno)

	roblox := login("otavi", "#lu1s")
	fmt.Println(roblox)
	loja := comprarGiftCard(20.00)
	fmt.Println(loja)

	for numero := 0; numero < 10; numero++ {
		pegarHoraAtual()
	}
	andar100Blocos(100)
}

func validar(nome string) bool {
	switch nome {
	case "otavio":
		return true
	default:
		return false
	}
}

func login(nome string, senha string) string {
	if nome == "otavio" && senha == "#lu1s" {
		return "logado"
	} else {
		return "nome ou senha incorreta"
	}
}

func pegarHoraAtual() {
	fmt.Println(time.Now())
}

// função chamada comprarGiftCard que receberá o valor que vc tem em carteira, e com switch validar, se:
// -- valor <= 10, retornar "voce pode comprar 10 itens"
// -- valor <= 20  retornar "voce pode comprar 20 itens"
// -- default "voce pode comprar muitos itens!"
func comprarGiftCard(carteira float64) string {
	switch carteira {
	case 10:
		return "voce pode comprar 10 itens"
	case 20:
		return "voce pode comprar 20 itens"
	default:
		return "voce pode comprar muitos itens!"
	}
}

// criar função andar100Blocos
// função deve receber o número de passos que deve andar
// se o numero de passos for igual a 100, deve fazer um for de 100 passos dentro deste if
// se não for igual a 100, deve printar "você pode andar mais!!"
func andar100Blocos(passos int) {
	if passos == 100 {
		for numero := 0; numero < 101; numero++ {
			fmt.Println(numero)
		}

	} else {
		fmt.Println("você pode andar mais!!")
	}
}
