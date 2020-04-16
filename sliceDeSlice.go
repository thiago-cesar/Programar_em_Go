/*sliceDeSlice.go*/

package main

import (
	"fmt"
	"strings"
)

func main(){

	tabuleiro:=[][] string{
		[]string{"X", "_", "_"},
		[]string{"0", "X", "0"},
		[]string{"_","0","X"},
	}

	

	//fmt.Println(tabuleiro[2])

	for i:=0; i< len(tabuleiro); i++{ //Percorrer cada linha com quebra de linha.
		//fmt.Println(tabuleiro[i])//Ou usar a função "join".
		fmt.Print(" \n", strings.Join(tabuleiro[i]," "))
	}
	
}