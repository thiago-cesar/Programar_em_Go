/*fnvalues3.go*/

package main

import "fmt"

func cumprimentar(nome string)func() string {
	//return "Ola, " + nome + "!"

	//return fmt.Sprintf("Olá, %s !", nome)

	return func () string{
		return fmt.Sprintf("Olá, %s !", nome)
	}
}


func main(){

	cumprimentarJoao := cumprimentar("João")

	fmt.Println(cumprimentarJoao())


}
