package main

import (
	"fmt"
	dz3Lesson "github.com/Nurbek11/dz3Lesson/fibonacci"
	dz3Lesson2 "github.com/Nurbek11/dz3Lesson/reverse"
	dz3Lesson3 "github.com/Nurbek11/dz3Lesson/runeByIndex"
	dz3Lesson4 "github.com/Nurbek11/dz3Lesson/sortImports"
	"strconv"
)

func main() {

	var word string = "OneTech"
	var idx int = 500
	fmt.Println("ATOI")
	fmt.Println(strconv.Atoi("15"))

	fmt.Println("ITOA:")
	fmt.Println(strconv.Itoa(idx))

	fmt.Println("REVERSE:")
	fmt.Println(dz3Lesson2.Reverse(word))

	fmt.Println("FIBO:")
	generator := dz3Lesson.Fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Print(generator(), " ")
	}

	fmt.Println("RUNE BY INDEX:")
	fmt.Println(dz3Lesson3.RuneByIndex(&word, &idx))

	//SORT IMPORTS
	dz3Lesson4.Sort("main.go")
}
