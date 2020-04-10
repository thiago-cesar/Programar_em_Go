/* variaveis.go*/

package main

import "fmt"


func main(){
	/*var idade, altura int = 35, 169 /* variáveis do mesmo tipo na mesma linha. Separação de valor repsectivamente.*/
	/*var nome string ="Thiago"*/
	/* ou declarar dessa forma e por inferência de tipo:
	var(
		idade = 35
		altura = 169
		nome = "Thiago" /*Se não houver uma variável, deve-se declarar o tipo
	)*/

	/*Podemos declarar também da forma abaixo:*/

	idade := 35
	altura:= 169
	nome:= "Thiago"

	fmt.Println("Meu nome é ", nome , " e tenho ", idade, " anos.")
	fmt.Println("Minha altura é: ",altura)
}

