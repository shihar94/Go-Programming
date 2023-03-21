package main

import "fmt"

var val string = "hello"


func main(){

	val := 42
	fmt.Println("local val:" , val)
	printGlobalVal()
	updateGlobalVal()
	

	*(&val) = 43
	fmt.Println(val)


}

func printGlobalVal(){
	fmt.Printf("global val: %v %T\n",val,val)
}

func updateGlobalVal(){
	val = "updated hello"
	printGlobalVal()
}