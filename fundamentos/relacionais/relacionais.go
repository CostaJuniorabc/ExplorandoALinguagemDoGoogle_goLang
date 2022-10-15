package main

import (
	"fmt"
	"time"
)

func main() {

	/* operadores relacionais
	== igual
	!= diferente
	< menor que
	> maior que
	<= menor igual
	>= maior igual
	*/
	fmt.Println(" == Strings Iguais : ", "Banana" == "Banana")
	fmt.Println(" != Diferentes : ", 3 != 2)
	fmt.Println("< Menor que : ", 3 < 2)
	fmt.Println(" > Maior que : ", 3 > 2)
	fmt.Println("<= Menor igual a : ", 3 <= 2)
	fmt.Println(">= Maior igua a : ", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	d3 := time.Unix(1, 1)
	d4 := time.Unix(2, 2)

	fmt.Println("Verificar datas : ", d1 == d2)
	fmt.Println("Verificar Datas : ", d3.Equal(d4))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	p3 := Pessoa{"Carlos"}
	p4 := Pessoa{"Antonio"}
	fmt.Println("Pessoa : ", p1 == p2)
	fmt.Println("Pessoas : ", p3.Nome)
	fmt.Println("Pessoas : ", p4)

}
