// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println(fatorial(4))
}

func fatorial(x int) int {
	var resultado int

	if x == 0 || x == 1 {
		resultado = 1
	} else {
		resultado = x * fatorial(x-1)
	}

	return resultado
}
