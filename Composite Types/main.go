package main 

import (
	"fmt"
)

//structs
type author struct{
	first string
	last string 

}

func (a author) fullname() string{
	return a.first + " " + a.last
}

func (a *author) changeName (first , last string){
	 a.first  = first 
	 a.last   = last
}

type person struct{
	author
	penName string
}

func (p person) fullname() string{
	return fmt.Sprintf("%s (%s)" , p.author.fullname() , p.penName)
}

type authors struct{
	name string
}

func main(){

	a := author{
		first: "james",
		last: "neesham",
	}

	fmt.Printf("%#v\n",a)
	fmt.Println(a.fullname())
	a.changeName("James", "Neesham")
	fmt.Println(a.fullname())


	p := person{
		author:author{
			first: "Mark",
			last : "Twain",
		},
		penName: "Harry",
	}

	fmt.Println(p.fullname())

	//arrays
	var arr[3]int 
	fmt.Println(arr)
	arr[0] = 1//cant resize arrays
	fmt.Println(arr)

	//slices manage with variable size arrays
	names := make([] string , 0) //created a string array with 0 size

	names = append(names, "Shihar")
	names = append(names, "Harry")
	names = append(names, "Ron")
	fmt.Println(names[0])
	fmt.Println(names[1:3])

	//maps 
	var authorList map[string]authors
	authorList = make(map[string] authors)
	authorList["mt"] = authors{name:"Mark Twain"}
	authorList["ws"] = authors{name:"William Shakespere"}

	fmt.Printf("%#v\n",authorList)


}