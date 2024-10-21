package main

import (
	"fmt"
)

func printStarLine(numOfStars int) string {
	if numOfStars <= 1 {
		return "*"
	} else {
		return "*" + printStarLine(numOfStars-1)
	}
}

func printStars(numOfStars int) {
	if numOfStars <= 1 {
		fmt.Println("*")
	} else {
		printStars(numOfStars - 1)
		fmt.Println(printStarLine(numOfStars))
	}
}

func main() {
	var num int
	fmt.Scanln(&num)
	printStars(num)
}
