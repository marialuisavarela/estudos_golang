package main

import (
	"fmt"
)

func main() {
	fmt.Println("------------- CALCULADORA -------------")
	switch getOperacao() {
	case "soma":
		var numerosSoma []float64
		var num float64
		for {
			fmt.Println("Digite o número que deseja somar. Digite '0' quando desejar parar")
			fmt.Scanln(&num)
			if num == 0 {
				break
			} else {
				numerosSoma = append(numerosSoma, num)
			}
		}
		fmt.Println("A soma é: ", soma(numerosSoma))

	case "subtração":
		var num1, num2 float64
		fmt.Println("Digite os números que deseja subtrair (com espaço)")
		fmt.Scanln(&num1, &num2)
		fmt.Println("A subtração é: ", subtracao(num1, num2))

	case "multiplicação":
		var numerosMulti []float64
		var num float64
		for {
			fmt.Println("Digite o número que deseja multiplicar. Digite '1' quando desejar parar")
			fmt.Scanln(&num)
			if num == 1 {
				break
			} else {
				numerosMulti = append(numerosMulti, num)
			}
		}
		fmt.Println("A multiplicação é: ", multiplicacao(numerosMulti))

	case "divisão":
		var n1, n2 float64
		fmt.Print("Digite os números que deseja dividir: ")
		fmt.Scanln(&n1, &n2)
		fmt.Printf("A divisão é: %.2f\n", divisao(n1, n2))

	default:
		fmt.Println("Você deve digitar uma opção válida: \nsoma \nsubtração \nmultiplicação \ndivisão \nequação1")
	}
}

func getOperacao() string {
	var operacao string
	fmt.Println("Digite a operação que deseja executar")
	fmt.Scanln(&operacao)
	return operacao
}

func soma(numeros []float64) float64 {
	var soma float64
	for _, v := range numeros {
		soma = soma + v
	}
	return soma
}

func subtracao(num1, num2 float64) float64 {
	return num1 - num2
}

func multiplicacao(numerosMulti []float64) float64 {
	multiplicacao := 1.0
	for _, v := range numerosMulti {
		multiplicacao = multiplicacao * v
	}
	return multiplicacao
}

func divisao(num1, num2 float64) float64 {
	return num1 / num2
}
