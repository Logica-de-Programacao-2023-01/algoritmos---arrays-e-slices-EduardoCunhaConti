package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var valor int
	fmt.Print("Diga um valor para ser procurado: ")
	fmt.Scan(&valor)

	found := false
	for _, numero := range array {
		if numero == valor {
			found = true
			break
		}
	}

	if found {
		fmt.Printf("O valor %d está presente no array.\n", valor)
	} else {
		fmt.Printf("O valor %d não está presente no array.\n", valor)
	}
}
