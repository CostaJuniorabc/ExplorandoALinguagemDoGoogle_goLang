package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0
	for i <= 200 { // não tem while em go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 400; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Print("Par : ")
		} else {
			fmt.Print("Impar : ")
		}
	}

	for {
		// laço infinito
		fmt.Println("Para Sempre... :(")
		time.Sleep(time.Second * 3)
	}

	// Veremos o foreach no capítulo de array
}
