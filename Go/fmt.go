package main

import "fmt"

func main() {
	age := 35
	name := "shaun"
	fmt.Print("Hello, ")
	fmt.Print("ninjas!\n")

	fmt.Println("Hello, ninjas!")
	fmt.Println("goodbye, ninjas!")
	fmt.Println("Hello,", "ninjas!")
	fmt.Println("my age is", age, "and my name is", name)

	fmt.Printf("my age is %v and my name is %v\n", age, name)
	fmt.Printf("my age is %q and my name is %q\n", age, name)
	//age does not work.

	fmt.Printf("age is of type %T\n", age)           //%T is for type
	fmt.Printf("you scored %f points!\n", 225.55)    //%0.1f is for float
	fmt.Printf("you scored %0.1f points!\n", 225.55) //%0.1f is for float

	var str = fmt.Sprintf("my age is %v and my name is %v\n", age, name) //returns the result as a string instead of printing it
	fmt.Println("the value stored in str is:", str)

}
