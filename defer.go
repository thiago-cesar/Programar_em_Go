/*defer.go*/

//Adiar a execução de uma função utilizando uma "defer".


package main

import (
	"fmt"
)

func main(){

	defer fmt.Println("Oi!") //é executada após as linhas posteriores e retorna para ecxcutar essa linha.
	fmt.Println("Tudo bem?")

	
	
}