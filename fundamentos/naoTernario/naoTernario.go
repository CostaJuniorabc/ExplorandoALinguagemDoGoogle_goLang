package main

import "fmt"

// Não tem operador ternário
func obterResultado(nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	} else {
		return "Reprovado"
	}
}

func main() {
	fmt.Println("Nota Final : ", obterResultado(6.2))
}
