/* forloop.go*/

package main

import "fmt"


func main(){

// somar números inteiros de 1 a 10.

	var soma int  = 0

	// for inicialização;condição; incremento

	for i := 1; i <= 10; i = i + 1{
		soma = soma + i

		fmt.Println (" valores de i:", i)
	}

	fmt.Println(soma)
}