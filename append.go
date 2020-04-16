/*append.go*/

//Adicionar elementos a pedaços.

package main

import "fmt"

func main(){
	 
	var nomes [] string 

	//slice valor "nil".

	var nomes2 = append(nomes, "João")//adicionar no pedaçono final e um novo pedaço.
	nomes2 = append(nomes2, "Maria")
	nomes2= append(nomes2, "Ana", "Mateus", "Rogerio")

	fmt.Println(nomes)
	fmt.Println(nomes2)
	fmt.Printf("len=%d; cap= %d", len(nomes2),
	 cap(nomes2))


}