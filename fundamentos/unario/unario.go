// https://www.udemy.com/course/curso-go/learn/lecture/8603196#overview

package main

import "fmt"

func main() {

	x := 1
	y := 2

	//apenas postfi
	x++ // x += 1 ou x = x + 1
	fmt.Println("Valor : ", x)

	y-- // y -= 1 ou y = y - 1
	fmt.Println("Valor : ", y)

	//fmt.Println(x == y--)

}
