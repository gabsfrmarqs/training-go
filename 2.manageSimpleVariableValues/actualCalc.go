package main

import (
	"fmt"
)

func sub(v1 float64, v2 float64) float64{
	return v1 - v2
}

func add(v1 float64, v2 float64) float64{
	final := v1 + v2
	return final
}

func div(v1 float64, v2 float64) float64{
	return v1 / v2
}

func mul(v1 float64, v2 float64) float64{
	return v1 * v2
}

func main(){
	var v1 float64 = 7.3
	var v2 float64 = 2.7

	fmt.Printf("Values for v1 and v2: %v and %v\n", v1, v2)

	fmt.Printf("%v\n",add(v1,v2))
	fmt.Printf("%v\n",sub(v1,v2))
	fmt.Printf("%v\n",div(v1,v2))
	fmt.Printf("%v\n",mul(v1,v2))




	
}

