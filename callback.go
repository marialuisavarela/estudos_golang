// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	sliceInicial := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(somenteImpares(soma, sliceInicial...))
}

func somenteImpares(f func(s ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 == 1 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}

func soma(s ...int) int {
	soma := 0
	for _, v := range s {
		soma = soma + v
	}
	return soma
}
