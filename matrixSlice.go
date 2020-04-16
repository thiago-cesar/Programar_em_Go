/* matrixSlice.go*/
/*PERCORRER CADA ELEMENTO DE CADA ELEMENTO DE LINHA*/

package main

import (
	"fmt"
)

func main(){

	tabuleiro:=[][] string{
		[]string{"X", "_", "_"},
		[]string{"0", "X", "0"},
		[]string{"_","0","X"},
	}

	for i := 0; i< len(tabuleiro); i++{ //Percorrer cada linha com quebra de linha.
		//tabuleiro[i] ACESSAR LINHA
		
		for j := 0; j < len(tabuleiro[i]); j++{
			//tabuleiro[i][j]
			fmt.Println(tabuleiro[i][j])


		}

	}
	
}