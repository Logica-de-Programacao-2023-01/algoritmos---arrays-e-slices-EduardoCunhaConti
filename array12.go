package main

import "fmt"

func main() {
	array := []int{3, 4, 9, 16, 21}
	fmt.Println(array)

	var multiplos []int
	for _, numero := range array {
		if numero%3 == 0 {
			multiplos = append(multiplos, numero)
		}
	}

	fmt.Println(multiplos)
}
