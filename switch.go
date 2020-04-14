/* switch.go*/

package main

import (
	"fmt"
)

func main(){

	var nome = "João"

	switch nome {
	 //Não necessita do Break
	case "Ana":

		fmt.Println(" É a Ana")

	case "João":

		fmt.Println(" É o João")

	default:

		fmt.Println("Não conheço")
	}
}