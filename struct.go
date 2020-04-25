/*struct.go*/

package main

import "fmt"

/*var x int = 40// Encapsular x e y.
* var y int = 67 

fmt.Println("x = %d y = %d \n", x , y)*/

func main(){

	type Posicao struct {

		X int
		Y int

	}

	pos := Posicao{40,67}
	pos.Y = 51
	
	
	fmt.Println(pos)

}