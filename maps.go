/*maps.go*/

package main

import "fmt"

func main(){

	var notas map[string]int
	notas = make(map[string]int) 

	//"Ana"---> 9
	//"Maria"---> 10

	notas["Ana"] = 9
	notas["Maria"] = 10

	fmt.Println(notas)

	//Acessar o valor individualmente:

	fmt.Println(notas["Maria"])

	//Determinar se "João" está no map.

	valor, existe := notas["João"]// Dessa forma o resultado é um valor booleano (True/False).

	if existe{
		fmt.Println(valor)
	}

	//fmt.Println(valor)
	//fmt.Println(existe)

}