package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	//tipos numericos inteiro
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é : ", reflect.TypeOf(32000)) //reflect.Type() em Golang é usada para obter o tipo de v

	//sem sinal (so positivos)... unit8 unit16 unit32 unit64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxFloat64
	fmt.Println("O valor maximo do int é ", i1)
	fmt.Println("O tipo i1 é :", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela unicode (int32)
	fmt.Println("O rume é : ", reflect.TypeOf(i2))
	fmt.Println(i2)

	//numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é : ", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é :", reflect.TypeOf(49.99)) //float64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é : ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string]
	s1 := "Olá meu nome é Junior"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é : ", len(s1))

	//string com multiplas linhas
	s2 := "Olá, meu nome é Jr"
	fmt.Println("O tamanho da string s2 é :", len(s2))

	//char ???
	//var x char ='b'
	char := 'a'
	fmt.Println("O tipo de char é : ", reflect.TypeOf(char))
	fmt.Println(char)

}
