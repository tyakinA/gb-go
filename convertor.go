package main

import (
	"fmt"
)

func showName(){
	var a int
	fmt.Print("введите сумму: ")
	fmt.Scan(&a)
	const b = 65
	var d = a / b

	fmt.Printf("Сумма в долларах: %v\n" ,d)
	  
}

func main (){
	showName()

 
}
