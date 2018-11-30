package main

import (
	"fmt"
)

func main() {
	// Esto es un comentario de una linea
	/*
		Esto es un comentario de mas de una
		linea
		:)
	*/
	fmt.Println("Variables! \n")

	// Declarar mas de una variable
	var indice, suma int = 1, 2
	mensaje := "Este tipo de declaración infiere el tipo"
	msg1, bool1, int1 := "Fantastico, no?", true, 8

	fmt.Printf("El indice es: %d y la suma esta en: %d\n", indice, suma)

	fmt.Printf("%s\n", mensaje)
	fmt.Println(msg1, " ", bool1)
	fmt.Printf("pool? Bola %d", int1)
}
