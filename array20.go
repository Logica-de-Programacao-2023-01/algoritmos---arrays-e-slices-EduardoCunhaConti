package main

import "fmt"

func main() {
	var n int
	fmt.Println("Diga o tamanho do array: ")
	fmt.Scanln(&n)

	fmt.Println("Diga o array de inteiros: ")
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	crescente := true
	for i := 1; i < n; i++ {
		if array[i] < array[i-1] {
			crescente = false
			break
		}
	}

	if crescente {
		fmt.Println("O array está em ordem crescente.")
	} else {
		fmt.Println("O array não está em ordem crescente.")
	}
}
