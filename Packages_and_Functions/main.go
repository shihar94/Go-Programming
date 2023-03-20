package main

//pkg.go.dev to explore other packages

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"github.com/jboursiquot/go-proverbs"
)

func greet() string{
	return "Hello"
}

//go doesnt support overloading
func greetWithName(name string) string{
	return "Hello" + name + "!"
}

func greetWithNameAndAge(name string , age int) (greeting string){
	greeting = "Hello, my name is " + name + " and i am " + strconv.Itoa(age) + " years old"
	return // this is called naked return as its specified greeting is the variable being returned
}

//multi value return 
func divide(a , b int) (int , error) {
	if b == 0 {
		return 0 , errors.New("Cannot divide by 0")
	}

	return a/b , nil
}

func getRandomProverbString() string{
	return proverbs.Random().Saying
}

func main(){
	const name , age = "Shihar" , 22
	fmt.Println(name , "is" , age ,"years old")
	fmt.Printf("Today is %s", time.Now().Weekday())
	fmt.Println(greetWithName("shihar"))
	fmt.Println(greetWithNameAndAge("Shihar", 29))
	fmt.Println(divide(10,5))
	fmt.Println(getRandomProverbString())
}