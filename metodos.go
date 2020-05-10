/*metodos.go*/

package main

import "fmt"

type Pessoa struct{
	Nome, Sobrenome string 	
}

//Método= oferecer uma funcionalidade a mais para a função.

func (p Pessoa) NomeCompleto() string {
	return p.Nome + " "+ p.Sobrenome

}

//O MÉTODO como uma função:

func FuncaoNomeCompleto(p Pessoa) string{
	return p.Nome + " "+ p.Sobrenome

}


func main(){

	alguem := Pessoa{"Jose", "de Alencar"} 

	fmt.Println(alguem.NomeCompleto())
	fmt.Println(FuncaoNomeCompleto(alguem))

}