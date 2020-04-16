/*slices2.go*/


package main

import (
	"fmt"
)
func main(){

	var nomes = [3] string {
		"Ana",
		"José",
		"Maria",
	}
	fmt.Println("Array inicial:", nomes)
	var p1[] string = nomes[0:2]
	fmt.Println(p1)
	p1[0] = "Rogerio"
	fmt.Println(p1)
	fmt.Println("Array original:", nomes)
}