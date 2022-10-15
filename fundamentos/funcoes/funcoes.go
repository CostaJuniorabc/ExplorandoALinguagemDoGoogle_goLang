package main

import "fmt"

//( função e um bloco de codigo que dentro dele tem uma sequencia de passo a ser executado )

// sunção para somar
func somar(parametro1 int, parametro2 int) int {
	return parametro1 + parametro2
}

// função para imprimir um valor
func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	resultado := somar(3, 5)
	imprimir(resultado)
}

// para rodar essa função tem que ser pelo termonal
//ancarlosda@ENGNB002678:~/Documentos/Junior/cursoGO/fundamentos/funçoes$
//comando: go run *.go
