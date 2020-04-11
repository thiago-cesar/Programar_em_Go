/* logica.go*/

package main

import "fmt"

func main(){
// && || !

// 3< x < 7
// x> 3 && x < 7

var x int = 4
var y int = 8

//Testar operador lógico .e:

fmt.Println("x> 3 E x < 7 =", x> 3 && x < 7)


//Testar operador logico .ou:


fmt.Println(" y < 3 OU y > 8 =",y < 3 || y > 8 )

// Operador de negação:

fmt.Println(" Negação de y = 8 =", y == 8)
}