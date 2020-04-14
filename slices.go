/*slices.go*/


package main

import (
	"fmt"
)
func main(){
						// 0  1  2  3  4  5  6
	var numeros = [7] int {1, 1, 2, 3, 5, 8, 13} // sem numeros ira aparecer 07 zeros

	//Slices: numeros[2,3,5,8]

	fmt.Println(numeros[2:])

	// numeros[2:]//omitindo o índice final- inicia do índice inicial e vai até o final.
	// numeros[:5]// omitindo o índice inicial e limitar o final ele inicia do ínício da array e vai até o final limite estipulado.
}
