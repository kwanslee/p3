package main

import "fmt"

func main() {

	fmt.Println("Hello, ninjas!")

	var nameOne string = "mario"
	//fmt.Println(nameOne) //not using the var will cause an error
	var nameTwo = "luigi" //type inference
	//fmt.Println(nameTwo)
	var nameThree string
	nameThree = "peach"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "bowser"
	nameThree = "yoshi"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "toad" //:= is shorthand for declaring and initializing a variable
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	var intOne int8 = 25
	//var intTwo int16 = 25
	//var intThree int32 = 25
	//var intFour int64 = 25
	var intTwo int8 = -128
	//var intThree uint=-25 //unsigned int cannot be negative. It will cause an error

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 883759234556096844566
	scoreThree := 1.5 //float64

	fmt.Println(ageOne, ageTwo, ageThree)
	fmt.Println(intOne, intTwo)
	fmt.Println(scoreOne, scoreTwo, scoreThree)
}

//go run main.go
