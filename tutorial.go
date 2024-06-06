package main

import "fmt"


func main() {

	var name string //implicit variable declaration, compilier will guess type and memory usage
	var number uint16 = 4000 // explict variable declaration for memory space allocation and data type
	
	 
	
	
	name = "dave" 
	fmt.Println(name,"\n",number)
	fmt.Printf("%T", name)	// "%T" is simalr to type fucntion in python
	number2 := 6 // walrus operator := can be used to skip type declaration
	//number2 = "hello" -> throws error, cannot redefine variable type
	fmt.Printf("%T", number2) 

	var bl bool //default value of false
	fmt.Println(bl)

}


