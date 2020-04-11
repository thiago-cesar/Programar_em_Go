/* funcoes2.go*/

package main

import "fmt"

/*func calcular(a int)(int, int){// segundo parênteses: tipo dos retornos.
	var quadrado int = a*a
	var cubo int = a*a*a
	return quadrado, cubo
}
*/

func calcular(a int)(quadrado int, cubo int){// segundo parênteses: tipo dos retornos.Pode declarar as variáveis de retorno, também.
	quadrado = a*a
	cubo =  a*a*a
	return quadrado, cubo
}

func main(){

	fmt.Println(calcular(2))


}