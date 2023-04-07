package main

import "fmt"

func main() {
	var n int
	fmt.Println("Diga o tamanho dos arrays: ")
	fmt.Scanln(&n)

	fmt.Println("Diga o primeiro array: ")
	array1 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array1[i])
	}

	fmt.Println("Diga o segundo array: ")
	array2 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array2[i])
	}

	var array3 []int
	for i := 0; i < n; i++ {
		array3 = append(array3, array1[i]+array2[i])
	}
	fmt.Println("A soma dos dois arrays é: ", array3)
}
