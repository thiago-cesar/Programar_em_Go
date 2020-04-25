/*struct2.go*/

package main

import "fmt"

func main(){
	type Posicao struct {
		X int
		Y int
	}

	var pos Posicao = Posicao{46, 87}//Sem valor, retorna o valor zeropara cada variável.
	fmt.Println(pos)
}