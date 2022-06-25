package main

import (
	"fmt"
)

// criar função chamada "venderProduto"
// função deve receber o nome do produto, e o valor do produto e devolver um bool dizendo se efetuou a venda ou não
// função deve imprimir o nome do produto, o valor do produto e se efetuou a venda ou não
// deve chamar a função "venderProduto" de dentro da função "main" e guardar o retorno da função em uma variável
// deve imprimir a variável que guardou o retorno da função "venderProduto"

// alterar metodo chamado venderPoduto para usar switch com as seguintes opções:
// valor do produto igual a 100
// valor do produto igual a 200
// valor do produto igual a 300
// só irá retornar true se o valor do produto for igual a 200
func main() {
	vendeuounao := venderProduto("iPhone 13", 3000.00)
	fmt.Println(vendeuounao)
}

func venderProduto(nomedoproduto string, valordoproduto float32) bool {
	fmt.Println(nomedoproduto)
	switch valordoproduto {
	case 100:
		fmt.Println("não vendeu")
		return false
	case 200:
		fmt.Printf("vendeu")
		return true
	case 300:
		fmt.Println("não vendeu")
		return false
	default:
		fmt.Println("não vendeu")
		return false
	}

	// if valordoproduto >= 2500.00 {
	// 	fmt.Println("vendeu")
	// 	return true
	// } else {
	// 	fmt.Println("não vendeu")
	// 	return false
	// }
}
