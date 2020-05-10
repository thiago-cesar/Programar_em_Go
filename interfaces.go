/*interfaces*/

package main

import (
	"fmt"
	"math"

)

type Geometrica interface {
	area() float64

}

//Interfaces é uma coletânea de várias assinaturas de função.

type Quadrado struct {

	lado float64 
}// area = lado*lado

func(q Quadrado) area() float64{
	return q.lado * q.lado
}

type Circulo struct {

	raio float64
}//area = pi* (raio* raio)

func(c Circulo)area() float64{
	return math.Pi* c.raio * c.raio

}


func main(){
	var g Geometrica
	g = Quadrado{3}

	fmt.Printf("a área do quadrado é %v\n", g.area())

	g = Circulo{5}
	fmt.Printf("A area do circulo é %v\n", g.area())

}