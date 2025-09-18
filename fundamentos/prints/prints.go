package main

import "fmt"

func main() {
	// nao adiciona quebra de linha
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// adiciona quebra de linha
	fmt.Println(" Nova")
	fmt.Println("Linha.")

	x := 3.141516
	// adiciona espaço entre os valores e concatena
	fmt.Println("O valor de x é", x)
	// controla a quantidade de casas decimais
	fmt.Printf("O valor de x é %.2f\n", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
}
