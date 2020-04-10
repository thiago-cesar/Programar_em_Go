/* coercao.go*/

package main

import "fmt"

func main(){
	//var a int = 46
	var b float64 = 6.4

	// var c float64 = a /* Go não converte inteiro para float*/
	

	// Conversão de float em inteiro:

	var c int = int(b)

	fmt.Println("O valor de c é: ",c)

}
