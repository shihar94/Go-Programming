package main

import (
	"fmt"
	"unicode"
)

//declaring package level variables
//compiler will assign zero as default value
var d , e , f int 

//go infers the type info from the values we assign 
var x , y , z = true , false , true

const pi = 3.15159
const a rune = 'a'

func main(){
	fmt.Println("d , e , f:",d , e, f)

	d = 1000000
	fmt.Printf("d: %d\n",d)

	//the following are not interefered with variables
	//from package level and do not need var keyword
	x , y , z := 0 , 1.25 , "hello"
	fmt.Println("x,y,z:",x,y,z)

	printVars()

	fmt.Printf("pi: %v -%T\n",pi,pi)
	fmt.Printf("a:  %v  -%T\n",a,a)

	unicode.IsLetter(a)

	c := 42.5 
	d := uint(c)
	fmt.Printf("%v=%T , %v=%T\n" , c,c,d,d)
	
	//pointers
	var f *int 
	g := 12
	f = &g
	fmt.Println(f) // prints the address currently held in f
	fmt.Println(*f) //to print the actual value
	

}

//func which uses package level variables

func printVars(){
	fmt.Println("x,y,z:",x,y,z)
}