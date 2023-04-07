package main

import "fmt"

func main() {
	array := [6]float64{1.2, 3.4, 5.6, 7.8, 9.0, 10.1}
	fmt.Println(array)

	var numero float64
	fmt.Println("Diga um número para ser adicionado a todos os elementos do array: ")
	fmt.Scan(&numero)

	for i := 0; i < len(array); i++ {
		array[i] += numero
	}

	fmt.Println("Array resultante:", array)
}
