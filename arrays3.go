/*arrays3.go*/

package main

import (
	"fmt"
)
func main(){

//Arrays
// 1, 1, 2, 3, 5, 8, 13

//Inserir os dados de uma vez.Inserir o operador de "=".

	var numeros = [7] int {1, 1, 2, 3, 5, 8, 13} // sem numeros ira aparecer 07 zeros
	// var numeros = [] int ...Colchetes vazios deduz-se que a quantidade elementos será o que foi inserido.
	//Acessar um determinado elemento via índice.
	//fmt.Println(numeros[1])  

	//Acessar via loop:

	for indice := 0; indice < len(numeros);indice++{
		fmt.Println(numeros[indice]) 
	}
}