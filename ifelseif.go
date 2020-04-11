/* ifelseif.go */

package main

import "fmt"


func main(){

// if else ifelse

	var a int = 4 // A variável pode ser definida na linha do if (apenas de dentro de seu escopo).

	if a > 10 {

	fmt.Println("O valor é maior que 10!")
	
	} else if a > 5 {
		fmt.Println(" O valor é maior que 5 e manor que 10. ")

	} else {
		fmt.Println("O valor é menor que 5.")
	}

}