/* whileloop.go*/

package main

import "fmt"
import "time" // utilizar para utilizar a função sleep
func main(){

	//somar:2+4+8+16+...

	// 2^1+2^2+2^3....

	var soma int = 2

	// for soma < 600 { // uma espécie de while

	for true{

	
		soma += soma

		fmt.Println("O valor da soma whileloop é:", soma)
	
		time.Sleep(100* time.Millisecond)
	}



}
