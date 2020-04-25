/*ponteiros*/

//Ponteiros são endereços de memória

package main

import "fmt"

func main(){

	a:=32

	p := &a
	
	fmt.Println(*p)

	//Mudar o valor de a para 18(exemplo):

	*p = 18

	fmt.Println(a)

	//Adicionar um valor ao ponteiro:
	*p = *p + 2

	fmt.Println(a)

}
