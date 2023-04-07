package main

import "fmt"

func main() {
	slice := []int{4, 8, 1, 5, 10, 3, 9, 7, 6, 2}

	var maior, menor int = slice[0], slice[0]
	for _, numero := range slice {
		if numero > maior {
			maior = numero
		}
		if numero < menor {
			menor = numero
		}
	}

	fmt.Println("O maior e o menor valor são:", maior, menor)
}
