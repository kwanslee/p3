package main

import "fmt"

func main() {
	//var ages [3]int = [3]int{20, 25, 30}
	//copy the left hand side to the right hand side.
	var ages = [3]int{20, 25, 30} //type inference
	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	//slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
}
