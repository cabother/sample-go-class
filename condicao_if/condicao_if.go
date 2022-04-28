package main

import "fmt"

func main() {
	result := login("otavio", "123abc")
	fmt.Println(result)

	fmt.Println("Hello World - IF")
}

// Métodos e condições IF's / ELSE
// - Criar um método chamado Login que receberá um nome de usuário e uma senha
// - Método deve retornar uma string
// - Usando um IF para validar e um ELSE para quando não for verdadeira a comparação:
//     - Se o usuário for "otavio" e a senha for "123abc" retornar "Usuário logado com sucesso"
//     - Se usuário ou senha diferentes dos citados acima, retornar "Usuário ou senha incorreto"
func login(nome string, senha string) string {
	if nome != "otavio" && senha != "123abc" {
		return "usuario e senha incorretos"
	}
	if nome != "otavio" {
		return "Usuário incorreto"
	}
	if senha != "123abc" {
		return "senha incorreta"
	} else {
		return "Login efetuado, vamos começar o jogo!"
	}
}
