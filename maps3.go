
/*maps3.go*/ 


package main

import "fmt"



type Perfil struct {

	Altura float64
	Ativo bool
	Profissão string}

func main(){

	var perfis map[string]Perfil = make(map[string]Perfil)

	perfis["João"] = Perfil {
		1.74, true, "Médico"}

	perfis["Maria"] = Perfil{
		1.63, false, "Engenheira",
	}

	fmt.Println(perfis)	

}