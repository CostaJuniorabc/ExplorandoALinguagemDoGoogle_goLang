// operadores aritimeticos
package main

import (
	"fmt"
	"math"
)

func main() {

	a := 3
	b := 2
	c := 3.0
	d := 2.0

	//funçoes aritimeticas
	fmt.Println("Funções Aritimeticas")
	fmt.Println("Resultado Soma => ", a+b)
	fmt.Println("Resultado Subtração => ", a-b)
	fmt.Println("Resultado Divisão => ", a/b)
	fmt.Println("Resultado Multiplicação => ", a*b)
	fmt.Println("Resultado Módulo => ", a%b)
	fmt.Println()

	//bitwise
	fmt.Println("bitwise")
	fmt.Println("E => ", a&b)   // 11 & 10 = 10
	fmt.Println("OU => ", a|b)  // 11 | 10 =11
	fmt.Println("Xor => ", a^b) // 11 ^ 10 = 01

	//outras operações usando math...
	fmt.Println("Maior => ", math.Max(float64(a), float64(b)))
	fmt.Println("Menor => ", math.Min(c, d))
	fmt.Println("Expodenciação => ", math.Pow(c, d))

}
