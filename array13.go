package main

import "fmt"

func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	var numero int
	fmt.Println("Diga um número para ser adicionado ao primeiro e ao último número do array: ")
	fmt.Scan(&numero)

	array[0] += numero
	array[6] += numero
	fmt.Println(array)
}
