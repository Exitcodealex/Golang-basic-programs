package main

import "fmt"

func main() {
	var i, j int
	var A = true
	var B = true
	var C = false
	var D = false
	fmt.Println("Give " + "first " + "number")
	fmt.Scan(&i)
	fmt.Println("Give " + "second " + "number")
	fmt.Scan(&j)
	fmt.Println("The sum of the 2 numbers is:", i+j)
	fmt.Println("The division result of the numbers is:", i/j)
	fmt.Println((A && B), (A || B), (A || C), (C || D))
	fmt.Println(!(A && B), !(A || B), !(A || C), !(C || D))
}
