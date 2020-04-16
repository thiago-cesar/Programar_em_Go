/*slices3.go*/


package main

import "fmt"

func main(){

	var nomes = [] string {
		"Ana",
		"José",
		"Maria",
	}

	//nomes = nomes[:2]- capacidAade 3.
	nomes = nomes[1:] //capacidade 2.

	fmt.Println(nomes)
	fmt.Printf("len= %d , cap=%d ", len(nomes), cap(nomes))
	

}