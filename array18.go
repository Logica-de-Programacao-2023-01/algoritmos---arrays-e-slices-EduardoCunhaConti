package main

import "fmt"

func eprimo(numero int) bool {
	if numero < 2 {
		return false
	}
	for i := 2; i*i <= numero; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Println("Diga um número. ")
	fmt.Scanln(&n)

	var primo []int
	numero := 2
	for len(primo) < n {
		if eprimo(numero) {
			primo = append(primo, numero)
		}
		numero++
	}
	fmt.Println("Primeiros", n, "números primos:", primo)
}