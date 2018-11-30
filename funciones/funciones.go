package main

import (
	"fmt"
)

func suma(x int, y int) int {
	return x + y
}

func main() {
	var primerNumero = 2
	segundoNumero := 5

	resultado := suma(primerNumero, segundoNumero)

	fmt.Printf("El resultado es: %d\n", resultado)
}
