package main

import (
	"fmt"
	"math"
)

func main() {
	a := math.Pow(3, 2)
		
	b := math.Pow(4, 2)
		
	c := math.Sqrt(a + b)
	fmt.Printf("%.1f", c)
	
}