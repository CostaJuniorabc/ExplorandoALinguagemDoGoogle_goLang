package main

import "fmt"

func main() {

	i := 1

	//go nao tem aritimetica de ponteiro
	var p *int = nil
	p = &i // pegando o endereço da variavel
	*p++   // desreferenciando (pegando o valor)
	i++

	//fmt.Println("Endereço de menoria : %v, Valor associado no ponteiro : %f, variavel normal : %f, referencia do valor da variavel : %f", p, *p, i, &i)
	fmt.Println(p, *p, i, &i)

}
