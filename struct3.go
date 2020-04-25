/*struct3.go*/

//Ponteiros para struct.

package main

import "fmt"

func main(){
	type Posicao struct {
		X int
		Y int
	}

	var pos Posicao = Posicao{47, 81}
	var k *Posicao = &pos
	(*k).Y = 33 //ou sem asterisco k.Y = 33.

	fmt.Println(	(*k).Y	)
}