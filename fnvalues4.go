/*fnvalues4.go*/

package main

import "fmt"

func adicionador() func (int) int{

	var soma int = 0
//closeura de uma função
	return func(a int)int {
		soma = soma + a
		return soma
	}
}


func main(){

	var ad1 = adicionador()
	fmt.Println(ad1(21))
	fmt.Println(ad1(13))//guardar ovalor anterior e soma a esse.

	
}

