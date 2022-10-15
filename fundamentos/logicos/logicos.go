package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 // ou exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete

}

func main() {

	tv50, tv32, sorvete := compras(true, true)
	fmt.Println("Comprar Tv de 50 : ", tv50)
	fmt.Println("Comprar Tv de 32 : ", tv32)
	fmt.Println("Comprar Sorrvete : ", sorvete)
	fmt.Println("Sorvete e Saudavel : ", !sorvete)
	fmt.Println()
	fmt.Printf("Comprar Tv de 50 : %t, Comprar Tv de 32 : %t, Comprar Sorrvete : %t, Sorvete e Saudavel : %t", tv50, tv32, sorvete, !sorvete)

}
