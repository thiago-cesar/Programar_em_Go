/*fnvalues2.go*/

package main

import "fmt"

//Função como argumento de outra função.

func computar(fn func(float64, float64)float64)float64{
	return fn(5, 6)
}


func main(){
	//Aramzenar uma função em uma variável.

	somar := func(a , b float64 ) float64{

		return a + b
	}

	/*funcaoMultiplicar := func(a, b float64)float64{
		return a * b
	}*/

	fmt.Println(computar(somar))

}