/*maps4.go*/

package main

import "fmt"

type Perfil struct {

	Altura float64
	Ativo bool
	Profissão string}

func main(){
	var perfis map[string]Perfil = map[string]Perfil{
		"João": Perfil{
		1.74, true, "Medico",
		},

		"Maria": Perfil{
		1.63, false, "Engenheira",
		},

	}


	fmt.Println(perfis)
}