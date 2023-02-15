package main

import "fmt"

func main() {
	switch getOperacao() {
	case "soma":
		fmt.Println("A soma é: ", somaComFor())
		//somaComMake()
		//fmt.Println("A soma é: ", somaComQtd())
	case "subtração":
		fmt.Println("A subtração é: ", subtracao())
	case "multiplicação":
		fmt.Println("A multiplicação é: ", multiplicacao())
	case "divisão":
		fmt.Printf("A divisão é: %.2f", divisao()) // não sei pq fica com o símbolo de porcentagem
	default:
		fmt.Println("Você deve digitar uma opeção válida: \nsoma \nsubtração \nmultiplicação \ndivisão")
	}
}

func getOperacao() string {
	var operacao string
	fmt.Println("Digite a operação que deseja executar")
	fmt.Scanln(&operacao)
	return operacao
}

func somaComQtd() float64 {
	var qtdNumeros int
	fmt.Println("Digite quantos números deseja somar")
	fmt.Scanln(&qtdNumeros)
	var numeros []float64

	var num float64
	for i := 0; i < qtdNumeros; i++ {
		fmt.Println("Digite o número que deseja somar")
		fmt.Scanln(&num)
		numeros = append(numeros, num)
	}

	return soma(numeros)
}

// func somaComMake() float64 {
// 	var qtdNumeros int
// 	fmt.Println("Digite quantos números deseja somar")
// 	fmt.Scanln(&qtdNumeros)
// 	numeros := make([]float64, 2)

// 	fmt.Println("Digite os números que deseja somar")
// 	fmt.Scanln(numeros...)

// 	return soma(numeros)
// }

func somaComFor() float64 {
	var numeros []float64
	var num float64
	for {
		fmt.Println("Digite o número que deseja somar. Digite '0' quando desejar parar")
		fmt.Scanln(&num)
		if num == 0 {
			break
		} else {
			numeros = append(numeros, num)
		}
	}

	return soma(numeros)
}

func soma(numeros []float64) float64 {
	var soma float64
	for _, v := range numeros {
		soma = soma + v
	}
	return soma
}

func subtracao() float64 {
	var num1, num2 float64
	fmt.Println("Digite os números que deseja subtrair (com espaço)")
	fmt.Scanln(&num1, &num2)
	return num1 - num2
}

func multiplicacao() float64 {
	var numeros []float64
	var num float64
	for {
		fmt.Println("Digite o número que deseja multiplicar. Digite '1' quando desejar parar")
		fmt.Scanln(&num)
		if num == 1 {
			break
		} else {
			numeros = append(numeros, num)
		}
	}

	multiplicacao := 1.0
	for _, v := range numeros {
		multiplicacao = multiplicacao * v
	}
	return multiplicacao
}

func divisao() float64 {
	var num1, num2 float64
	fmt.Print("Digite os números que deseja dividir: ")
	fmt.Scanln(&num1, &num2)
	return num1 / num2
}
