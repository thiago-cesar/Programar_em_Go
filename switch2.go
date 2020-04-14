/* switch2.go*/

package main

import (
	"fmt"
)

func main(){

	var nota int = 4

	switch true {// o true também pode ser omitido. O mesmo fica implícito.

	case nota >9:

		fmt.Println("Ótimo!")

	case nota>7:

		fmt.Println("Muito Bem!")

	case nota>6:

		fmt.Println("Bom!")

	default:

		fmt.Println("Péssimo!")
	
	}
}