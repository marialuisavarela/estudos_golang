package main

import (
	"fmt"
	"math"
)

func main() {
	exibirMenu()
	switch getOperacao() {
	case 1:
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

	case 2:
		var num1, num2 float64
		fmt.Println("Digite os números que deseja subtrair (com espaço)")
		fmt.Scanln(&num1, &num2)
		fmt.Println("A subtração é: ", subtracao(num1, num2))

	case 3:
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

	case 4:
		var n1, n2 float64
		fmt.Print("Digite os números que deseja dividir: ")
		fmt.Scanln(&n1, &n2)
		fmt.Printf("A divisão é: %.2f\n", divisao(n1, n2))

	case 5:
		var a, b float64
		fmt.Println("Equação de Primeiro Grau no formato de ax + b = 0")
		fmt.Println("Digite os valores de a e b, separados por espaço")
		fmt.Scanln(&a, &b)
		fmt.Println("O valor de x é: ", equacao1(a, b))
	case 6:
		var a, b, c float64
		fmt.Println("Equação de Segundo Grau no formato de ax² + bx + c = 0")
		fmt.Println("Digite os valores de a, b e c, separados por espaço")
		fmt.Scanln(&a, &b, &c)
		fmt.Println("Os valores de x são: ", equacao2(a, b, c))

	default:
		fmt.Println("Você deve digitar uma opção válida")
	}
}

func exibirMenu() {
	fmt.Println("------------- CALCULADORA -------------")
	fmt.Println("1 - Soma")
	fmt.Println("2 - Subtração")
	fmt.Println("3 - Multiplicação")
	fmt.Println("4 - Divisão")
	fmt.Println("5 - Equação de Primeiro Grau")
	fmt.Println("6 - Equação de Segundo Grau")
}

func getOperacao() int {
	var operacao int
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

func equacao1(a, b float64) float64 {
	resultado := -b / a

	return resultado
}

func equacao2(a, b, c float64) []float64 {
	var retorno []float64
	if a == 0 {
		fmt.Println("A equação não é de segundo grau")
		return retorno //feio
	} else {
		delta := math.Pow(b, 2) - 4*a*c
		if delta < 0 {
			fmt.Println("Não existe raiz real")
			return retorno
		} else {
			x1 := (-b + math.Sqrt(delta)) / (2 * a)
			x2 := (-b - math.Sqrt(delta)) / (2 * a)

			var resultado []float64
			resultado = append(resultado, x1, x2)

			return resultado
		}
	}
}
