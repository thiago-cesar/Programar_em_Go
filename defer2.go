/*defer2.go*/

package main

import "fmt"

func imprimir() string {
	fmt.Println("Imprimindo...")
	return "VALOR de IMPRIMIR"

}

func main(){

	defer fmt.Println(imprimir())
	fmt.Println("2")
	fmt.Println("3")


}