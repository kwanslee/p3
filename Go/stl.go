package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))         //true
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) //hi there friends!
	fmt.Println(strings.ToUpper(greeting))                   //HELLO THERE FRIENDS!

	fmt.Println(strings.Index(greeting, "ll")) //2
	fmt.Println(strings.Split(greeting, " "))  //[hello there friends!]

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages) //[20 25 30 35 45 50 60 75]

	index := sort.SearchInts(ages, 30) //2
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)                              //[bowser luigi mario peach yoshi]
	fmt.Println(sort.SearchStrings(names, "luigi")) //1

	//ages := map[string]int{"shaun": 23, "kim": 27}
}
