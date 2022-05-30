package main

import (
	"fmt"
)

// fazerfunc que se chama "comprarGiftCard"
// deve receber o valor que vc tem, e devolver o troco da compra de um produto de 15.99, se a compra for efetuada e devolver tbm + uma mensagem:
// se o valor pago menor que valor do produto (15.99) então a mensagem será saldo insuficiente
// se o valor for maior ou igual, voce efetuou a compra, mensagem "Compra finalizada com sucesso"

func main() {
	oRetorno, resto := comprarGiftCard(10.90, 15.99)
	fmt.Println(oRetorno)
	fmt.Println(resto)
}

func comprarGiftCard(meuDinheiro float64, custo float64) (string, float64) {
	if meuDinheiro >= custo {
		return "comprae fetuada", meuDinheiro - custo

	} else {
		return "saldo Insufisiente", meuDinheiro
	}

}
