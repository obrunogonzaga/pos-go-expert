package main

import (
	"fmt"
)

func main() {
	fmt.Println(soma(1, 3, 5, 7, 11, 18))
}

func soma(numeros ...int) (int) {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
