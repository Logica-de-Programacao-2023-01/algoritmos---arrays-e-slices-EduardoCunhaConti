package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array)

	var pares []int
	for _, numero := range array {
		if numero%2 == 0 {
			pares = append(pares, numero)
		}
	}
	fmt.Println(pares)
}
