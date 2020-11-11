package main

import (
	"fmt"
	"net"
	"encoding/gob"
	"time"
)



func proceso(numProceso *int, contador *int) {

	for{
		*contador++
		
		fmt.Println(*numProceso, ": ", *contador )
	
		time.Sleep(time.Millisecond * 500)

	}

	
}


func cliente(numProceso *int, contador *int){
	
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	nuevoCliente := true
	err = gob.NewEncoder(c).Encode(nuevoCliente)
	if err != nil {
		fmt.Println(err)
		
	 }

	
	err = gob.NewDecoder(c).Decode(numProceso)
	err2 := gob.NewDecoder(c).Decode(contador)
	


	if err != nil || err2 != nil {
		fmt.Println(err)
		fmt.Println(err2)
		return
	}else {

		go proceso(numProceso, contador)
		
		
	}	

	

	
	c.Close()
		
	

}


func regresarProceso(numProceso *int, contador *int){

	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gob.NewEncoder(c).Encode(*numProceso)
	if err != nil {
		fmt.Println(err)

	 }
	


}


func main(){

	var numProceso int
	var contador int
	
	go cliente(&numProceso, &contador)

	var input string
	fmt.Scanln(&input)

	//regresar el proceso
	regresarProceso(&numProceso, &contador)
	fmt.Println("Cliente terminado")


}