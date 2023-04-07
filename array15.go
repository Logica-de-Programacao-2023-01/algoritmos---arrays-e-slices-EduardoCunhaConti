package main

import "fmt"

func main() {
	array := [10]float64{1.2, 1.7, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9, 9.1}

	var maior []float64
	for _, numero := range array {
		if numero > 5 {
			maior = append(maior, numero)
		}
	}
	fmt.Println(maior)
}
