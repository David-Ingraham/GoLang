package main

import "fmt"

func main() {

	fmt.Printf("Hello %T %v", 10, 10) //%v value, %T type. %% prints literal percent sign

	var x string = fmt.Sprintf("Hello %T %v", 20, 20) //Sprintf saves a formated string to a variable
	fmt.Println("\n",x,"\n")	

	fmt.Printf("Hello %t", true)
	fmt.Println()

	fmt.Printf("Number 1024 in binary: %b", 1024)
	fmt.Println()
	fmt.Printf("Number 101 in decimal: %b", 101) //deosnt work like that. go take that second argumant as a base 10 int
	
}
