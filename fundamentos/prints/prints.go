package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print("Linha")

	//-----------------------------------------
	fmt.Println("")
	fmt.Println("Nova")
	fmt.Println("Linha")

	//----------------------------------------
	x := 3.141516

	fmt.Println("O valor de x é :", +x)

	//-------------------------------------------------
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é :" + xs)
	fmt.Println("O valor de x é :", x)

	fmt.Printf("O valor de x é %.2f", x) // pega so as 2 casas deimais %f = tipo float , %s = tipo string

	//--------------------------------------
	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %t %s %.1f", a, b, c, d, b) // /n usado para criar uma nova linnha sempre passar os valores corretos de % para o resultado vir corretamente
	fmt.Printf("\n%v %v %v %v", a, b, c, d)         // %v e um tipo generico para devolver valor
}
