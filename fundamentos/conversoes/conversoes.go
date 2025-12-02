package main

import (
	"fmt"
	c "strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Cuidado!
	fmt.Println("Teste " + string(rune(97)))

	// int para string
	fmt.Println("Teste " + c.Itoa(97))

	// string para int
	num, _ := c.Atoi("123")
	fmt.Println(num - 120)

	// string para bool
	b, _ := c.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
