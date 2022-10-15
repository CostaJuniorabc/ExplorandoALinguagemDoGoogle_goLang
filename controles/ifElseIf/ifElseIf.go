package main

import "fmt"

func notaParaConceito(n float64) string {
	if n > 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {

	fmt.Println("Nota Marcelo : ", notaParaConceito(9.8))
	fmt.Println("Nota Alaisa : ", notaParaConceito(8.8))
	fmt.Println("Nota Daniele : ", notaParaConceito(5.8))
	fmt.Println("Nota Junior : ", notaParaConceito(4.8))
	fmt.Println("Nota Lucas : ", notaParaConceito(1.8))

}
