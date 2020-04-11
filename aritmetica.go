/*aritmetica.go*/

package main

import "fmt"


func main(){

	//Operdadores:
	// + - * /
	// %

	var a int = 23
	var b float64 = 7

	fmt.Println(" a + b =", float64(a) + b) //variáveis com tipos diferentes ocasiona erro na compilação.
	fmt.Println("a - b =", float64(a) - b)
	fmt.Println("a * b =", float64(a) * b)
	fmt.Println("a / b =", float64(a) / b)

}