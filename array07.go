package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}

	var valor int
	fmt.Print("Diga um valor: ")
	fmt.Scan(&valor)

	found := false
	for _, numero := range array {
		if numero == valor {
			found = true
			break
		}
	}

	if found {
		fmt.Println(array)
	} else {
		array = append(array, valor)
		fmt.Println(array)
	}
}
