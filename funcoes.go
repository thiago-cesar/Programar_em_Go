/* funcoes.go*/

package main

import "fmt"

func soma( x int, y int) int { //Terceiro "int" correponde a declaração do tipo de variável de retorno.
								//Se todos os parâmetros da função forem do mesmo tipo, basta informar o type no final.
								
	return x+y

}

func main(){

	fmt.Println(soma(3,5))
	
}