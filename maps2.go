/*maps2.go*/
/*maps2*/ 


package main

import "fmt"

func main(){

	var notas map[string]int
	notas = map[string]int{
		"Ana":10,
		"Maria":9,
	}

	/*
	notas = make(map[string]int) 
	notas["Ana"] = 9
	notas["Maria"] = 10 
	*/

	fmt.Println(notas)	

}