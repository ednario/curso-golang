package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numeros inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal Inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("O byte e", reflect.TypeOf(b))

	// com sinal... int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 e", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	// por padrão sempre será float64
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O tipo bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Ola meu nome e Ednario"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string e", len(s1))
	s2 := `Olá
	meu
	nome
	e
	Ednario` // string raw (crua)
	fmt.Println("O tamanho da string s2 é", len(s2))

	// char (na verdade não existe um tipo char em Go)
	char := 'A' // representa um valor int32 (rune)
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}
