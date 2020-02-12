package main

import (
	"fmt"
)

func showName(){
	var a int
	fmt.Print("Введите сумму: ")
	fmt.Scan(&a)
	var b int
	fmt.Print("Ведите процент: ")
	fmt.Scan(&b)
	//var y = [...]int{1год, 2года, 3года, 4года, 5лет}
	for i := 0; i < 5; i++ {
		
		a = a + ( a * ( b / 100 ))
		
	}
}
func main (){
	showName("y [i] :" + a)
}
