package main

import "fmt"

func main() {
	matriz := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	var linha, coluna int
	fmt.Println("Diga uma linha e uma coluna: ")
	fmt.Scanln(&linha, &coluna)

	fmt.Printf("Valor na posição (%d, %d): %d\n", linha, coluna, matriz[linha][coluna])
}