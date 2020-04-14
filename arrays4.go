/*arrays4.go*/

package main

import (
	"fmt"
)
func main(){
	
	var numeros = [7] int {1, 1, 2, 3, 5, 8, 13} // sem numeros ira aparecer 07 zeros
	var soma int = 0

	for i := 0; i < len(numeros); i++{
		
		soma= soma +numeros[i]
		
	}
	fmt.Println(soma)

}