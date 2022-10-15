package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2
	nota := 6.9

	//converte floar para int
	notaFinal := int(nota)
	fmt.Println("Nota Final : ", notaFinal)

	//converter um tipo inteiro para um float
	fmt.Println(x / float64(y))

	//cuidado dentro vem o cod da tabela ask
	fmt.Println("Teste 1 = " + string(97))
	fmt.Println("Teste 2 = " + string(1234))

	// int para string
	fmt.Println("Conversão string = " + strconv.Itoa(123))

	//string para int (esse metodo 'strcov' retorna 2 coisas um numero ou um erro  caso passe um valor que nao seja uma string)
	num, _ := strconv.Atoi("123")
	fmt.Println("conversão para inteiro = ", num)
	fmt.Println("Resultado = ", (num - 122))

	b, _ := strconv.ParseBool("1")
	if b {
		fmt.Println("Verdadeiro")
	}

}
