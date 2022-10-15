package main

import (
	"fmt"
	"math"
)

func main() {

	const PI float64 = 3.1415 // constante nomedaconstante tipo e recebe valor
	var raio = 3.2            // tipo ( float64 ) inferido pelo compilador

	//forma redizina de criar uma variavel
	area := PI * math.Pow(raio, 2) // declarando e iniciando uma variavel

	fmt.Println("A área da circuferência é : ", area)

	// outras maneiras de declarar constantes
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println("Resultado : ", a, b, c, d)

	var e, f bool = true, false // declarar e atribuir valor
	fmt.Println("Resultado :", e, f)

	g, h, i := 2, false, "epa!!!" // declarar e atribuir valor

	fmt.Println("Resultado :", g, h, i)

	// obs sempre que criar uma variavel tem que utilizar a mesma
}
