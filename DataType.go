package main

import "fmt"

//global variable declaration
var (
	m int
	n int
)

func main () {
	var x int = 1 // integer data type
	var y int 	  // integer data type

	fmt.Println(x)
	fmt.Println(y)

	var a,b,c = 5.25,25.25,14.15 // Multiple float32 varible declaration
	fmt.Println(a,b,c)

	city := "Berlin" // String varible declaration
	Country := "Germany" // Variable names are case sensitive
	fmt.Println(city)
	fmt.Println(Country) // variable names are case senstive

	food,drink,price := "Pizza","Pepsi", 125  // Multiple type of variable declaration in same line
	fmt.Println(food,drink,price)
	m,n = 1,2
	fmt.Println(m,n)

}