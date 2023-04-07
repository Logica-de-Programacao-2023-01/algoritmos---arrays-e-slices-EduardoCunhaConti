package main

import "fmt"

func main() {
	slice := make([]string, 8)

	var termo string
	fmt.Print("Insira um valor para ser removido: ")
	fmt.Scan(&termo)

	slice = []string{"maçã", "banana", "morango", "banana", "morango", "cereja", "abacaxi", "banana"}

	resultado := make([]string, 0, len(slice))
	for _, palavra := range slice {
		if palavra != termo {
			resultado = append(resultado, palavra)
		}
	}

	fmt.Printf("Slice resultante: %v\n", resultado)
}
