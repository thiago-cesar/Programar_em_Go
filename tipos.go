/*tipos.go*/


/*int int8 int16 int32 int64

uint uint8 uint16 uint32 uint64

string

float32 float64

bool  true/false

complex64 complex128    a+bi


byte, rune*/

package main

import "fmt"


func main(){
	var altura = 1.69
	var aberto = true

	fmt.Printf("Tipo: %T \n Valor: %v \n", altura, altura)
	fmt.Printf("Tipo: %T Valor:%v", aberto,aberto)
}