package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}
	// for loop without ()
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"yoshi", "mario", "peach", "bowser"}

	for i := 0; i < len(names); i++ {
		fmt.Println("value of i is:", names[i])
	}

	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}
}
