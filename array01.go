package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	fmt.Println(array)

	soma := 0
	for i := 0; i < len(array); i++ {
		soma += array[i]
	}

	fmt.Println("A soma dos valores do array é", soma)
}
