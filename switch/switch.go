package main

import "fmt"

func main() {
	numero := 30

	switch numero {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("nothing")
	}

	nome := "otavio"
	switch nome {
	case "renato":
		fmt.Println("one")
	case "otavio":
		fmt.Println("two")
	case "joao":
		fmt.Println("three")
	default:
		fmt.Println("nothing")
	}

	numeroOtavio := 200
	switch numeroOtavio {
	case 1000:
		fmt.Println("one thousen")
	case 100:
		fmt.Println("a hundred")
	case 200:
		fmt.Println("two huudred")
	default:
		fmt.Println("default")
	}
}
