package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	switch getOperacao() {
// 	case "soma":
// 		fmt.Println("A soma é: ", somaComFor())
// 		//fmt.Println("A soma é: ", somaComQtd())
// 	case "subtração":
// 		fmt.Println("A subtração é: ", subtracao())
// 	case "multiplicação":
// 		fmt.Println("A multiplicação é: ", multiplicacao())
// 	case "divisão":
// 		fmt.Printf("A divisão é: %.2f\n", divisao())
// 	case "equação1":
// 		fmt.Println("O valor de x é: ", equacao1())
// 	case "equação2":
// 		fmt.Println("Os valores de x são: ", equacao2())
// 	default:
// 		fmt.Println("Você deve digitar uma opção válida: \nsoma \nsubtração \nmultiplicação \ndivisão \nequação1")
// 	}
// }

// func getOperacao() string {
// 	var operacao string
// 	fmt.Println("Digite a operação que deseja executar")
// 	fmt.Scanln(&operacao)
// 	return operacao
// }

// func somaComQtd() float64 {
// 	var qtdNumeros int
// 	fmt.Println("Digite quantos números deseja somar")
// 	fmt.Scanln(&qtdNumeros)
// 	var numeros []float64

// 	var num float64
// 	for i := 0; i < qtdNumeros; i++ {
// 		fmt.Println("Digite o número que deseja somar")
// 		fmt.Scanln(&num)
// 		numeros = append(numeros, num)
// 	}

// 	return soma(numeros)
// }

// func somaComFor() float64 {
// 	var numeros []float64
// 	var num float64
// 	for {
// 		fmt.Println("Digite o número que deseja somar. Digite '0' quando desejar parar")
// 		fmt.Scanln(&num)
// 		if num == 0 {
// 			break
// 		} else {
// 			numeros = append(numeros, num)
// 		}
// 	}

// 	return soma(numeros)
// }

// func soma(numeros []float64) float64 {
// 	var soma float64
// 	for _, v := range numeros {
// 		soma = soma + v
// 	}
// 	return soma
// }

// func subtracao() float64 {
// 	var num1, num2 float64
// 	fmt.Println("Digite os números que deseja subtrair (com espaço)")
// 	fmt.Scanln(&num1, &num2)
// 	return num1 - num2
// }

// func multiplicacao() float64 {
// 	var numeros []float64
// 	var num float64
// 	for {
// 		fmt.Println("Digite o número que deseja multiplicar. Digite '1' quando desejar parar")
// 		fmt.Scanln(&num)
// 		if num == 1 {
// 			break
// 		} else {
// 			numeros = append(numeros, num)
// 		}
// 	}

// 	multiplicacao := 1.0
// 	for _, v := range numeros {
// 		multiplicacao = multiplicacao * v
// 	}
// 	return multiplicacao
// }

// func divisao() float64 {
// 	var num1, num2 float64
// 	fmt.Print("Digite os números que deseja dividir: ")
// 	fmt.Scanln(&num1, &num2)
// 	return num1 / num2
// }

// func equacao1() float64 {
// 	var a, b float64
// 	fmt.Println("Equação de Primeiro Grau no formato de ax + b = 0")
// 	fmt.Println("Digite os valores de a e b, separados por espaço")
// 	fmt.Scanln(&a, &b)

// 	resultado := -b / a

// 	return resultado
// }

// func equacao2() []float64 {
// 	var a, b, c float64
// 	fmt.Println("Equação de Segundo Grau no formato de ax² + bx + c = 0")
// 	fmt.Println("Digite os valores de a, b e c, separados por espaço")
// 	fmt.Scanln(&a, &b, &c)
// 	var retorno []float64
// 	if a == 0 {
// 		fmt.Println("A equação não é de segundo grau")
// 		return retorno //feio
// 	} else {
// 		delta := math.Pow(b, 2) - 4*a*c
// 		if delta < 0 {
// 			fmt.Println("Não existe raiz real")
// 			return retorno
// 		} else {
// 			x1 := (-b + math.Sqrt(delta)) / 2 * a
// 			x2 := (-b - math.Sqrt(delta)) / 2 * a

// 			var resultado []float64
// 			resultado = append(resultado, x1, x2)

// 			return resultado
// 		}

// 	}
// }
