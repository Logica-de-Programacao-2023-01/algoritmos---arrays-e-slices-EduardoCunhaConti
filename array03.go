package main

import "fmt"

func main() {
	numeros := []float64{1.5, 2.5, 3.5, 4.5}
	fmt.Println(numeros)

	produto := 1.0
	for i := 0; i < len(numeros); i++ {
		produto *= numeros[i]
	}

	fmt.Println("O produto dos valores do array é", produto)
}
