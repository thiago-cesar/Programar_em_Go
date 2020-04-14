/* arrays2.go*/

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

	fmt.Println(numeros)  

}