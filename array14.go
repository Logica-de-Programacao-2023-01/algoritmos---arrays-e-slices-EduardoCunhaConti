package main

import "fmt"

func main() {
	slice := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice)

	var termo1, termo2 int
	fmt.Println("Diga dois índices para serem trocados de posição no slice: ")
	fmt.Scanln(&termo1, &termo2)

	slice[termo1], slice[termo2] = slice[termo2], slice[termo1]
	fmt.Println(slice)
}
