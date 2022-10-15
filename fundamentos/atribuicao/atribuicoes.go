package main

import "fmt"

func main() {

	var b byte = 3
	fmt.Println(b)

	i := 3 //inferência de tipo
	i += 3 // aditiva i = i + 3
	i -= 3 // subtrativa i = i - 3
	i /= 2 // divisiva i = i / 2
	i *= 2 // multiplicativa i = i x 2
	i %= 2 // modulo i = i % 2

	fmt.Println(i)

	x, y := 1, 2 //atribuir em 2 variaveis  x=1 y=2
	fmt.Println(x, y)

	x, y = y, x // atribui o valor de x=y y=x atribuidos quando criados anteriomente so inverteram os valores
	println(x, y)

}
