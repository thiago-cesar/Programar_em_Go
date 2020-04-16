/*slices4.go*/


package main

import "fmt"

func main(){

	var animes [3]string;
	animes = [3]string{"Drangon","Sailor","Yuyu"}

	var doisPrimeiros = animes[:2]
	var doisUltimos = animes[1:]

	fmt.Println(animes)
	fmt.Println(doisPrimeiros)
	fmt.Printf("len= %d  cap=%d", len(doisPrimeiros), 
	cap(doisPrimeiros))

	fmt.Println(animes)
	fmt.Println(doisUltimos)
	fmt.Printf("len= %d  cap=%d", len(doisUltimos), 
	cap(doisUltimos))


}