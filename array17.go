package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array)

	var numero int
	fmt.Println("Diga um número para ser adicionado aos índices pares do array: ")
	fmt.Scan(&numero)

	array[0] += numero
	array[2] += numero
	array[4] += numero
	array[6] += numero
	array[8] += numero
	fmt.Println(array)
}