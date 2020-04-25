/*range.go*/

package main

import "fmt"

func main(){
	
	numeros:= []int {10, 20, 30, 40}

/*	for i := 0; i< len(numeros); i++ {
		fmt.Println(numeros[i])
	}*/

for indice, valor := range numeros{// Inserir o "_" caso não ueira imprimir o índice ou o valor.
	fmt.Println(indice, valor)

}


}