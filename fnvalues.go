/*fnvalues.go*/

package main

import "fmt"


func main(){
	//Aramzenar uma função em uma variável.

	funcaoSomar := func(a , b float64 ) float64{
		return a + b
	}

	funcaoMultiplicar := func(a, b float64)float64{
		return a * b
	}

	fmt.Println(funcaoSomar(3,4))
	fmt.Println(funcaoMultiplicar(3,4))

}