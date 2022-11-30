package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		var opcao int
		var valor1 float64
		var valor2 float64

		fmt.Println("Selecione uma das operações:\n1. +\n2. -\n3. *\n4. /\n0. sair\nnúmero da opção:")
		fmt.Scan(&opcao)
		switch opcao {
		case 1:
			inputValores(&valor1, &valor2)
			fmt.Println("o resultado da soma é: ", (valor1 + valor2))
			valor1 = 0
			valor2 = 0
		case 2:
			inputValores(&valor1, &valor2)
			fmt.Println("o resultado da subtração é: ", (valor1 - valor2))
			valor1 = 0
			valor2 = 0
		case 3:
			inputValores(&valor1, &valor2)
			fmt.Println("o resultado da multiplicação é: ", (valor1 * valor2))
			valor1 = 0
			valor2 = 0
		case 4:
			inputValores(&valor1, &valor2)
			fmt.Println("o resultado da divisão é: ", (valor1 / valor2))
			valor1 = 0
			valor2 = 0
		case 0:
			fmt.Println("Saindo ...")
			os.Exit(0)
		}
	}
}

func inputValores(valor1, valor2 *float64) {
	fmt.Println("digite o valor 1: ")
	fmt.Scan(valor1)
	fmt.Println("digite o valor 2: ")
	fmt.Scan(valor2)
}
