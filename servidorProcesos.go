package main

import (
	"fmt"
	"net"
	"encoding/gob"
	"time"
)


func procesos(procesos *[]int, contador *int) {

	for{
		*contador++

		for _, p := range *procesos{

		fmt.Println(p, ": ", *contador )

		}
		fmt.Println("---------------------")

		
		time.Sleep(time.Millisecond * 500)

	}

	
}



func servidor (procesos *[]int, contador *int) {
	
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	
		for{
		c, err:= s.Accept()

		
		if err != nil{
			fmt.Println(err)
			continue
		}

		go handleClient(c, procesos, contador)


	}


		

	
}



func handleClient(c net.Conn, procesos *[]int, contador *int){

	var nuevoProceso int

	err := gob.NewDecoder(c).Decode(&nuevoProceso)
	if err != nil {
		// fmt.Println(err)
		
	}else {
		*procesos = append(*procesos, nuevoProceso )
		return
		
	}
	
	
	procesosNuevos := *procesos
	numProceso :=  procesosNuevos[0]
	*procesos = append(procesosNuevos[:0], procesosNuevos[0+1:]... )


	err = gob.NewEncoder(c).Encode(numProceso)
	if err != nil {
		fmt.Println(err)

	 }
	err = gob.NewEncoder(c).Encode(contador)
	if err != nil {
		fmt.Println(err)

	 }


	
		

	
}


func main(){

	
	contador := 0
	procesosSlice := []int{0, 1, 2, 3, 4}


	go procesos(&procesosSlice, &contador)
	go servidor(&procesosSlice, &contador)

	var input string
	fmt.Scanln(&input)

}