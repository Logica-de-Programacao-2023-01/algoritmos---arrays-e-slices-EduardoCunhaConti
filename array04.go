package main

import "fmt"

func main() {
	fmt.Print("Qual o tamanho do slice? ")
	var tamanho int
	fmt.Scan(&tamanho)

	slice := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Printf("Insira o valor do elemento %d: ", i)
		var valor int
		fmt.Scan(&valor)
		slice[i] = valor
	}

	fmt.Println("O slice é:", slice)
	soma := 0
	for i := 0; i < len(slice); i++ {
		soma += slice[i]
	}
	fmt.Println("A soma dos valores do slice é:", soma)
}
