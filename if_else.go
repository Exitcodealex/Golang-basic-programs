package main

import "fmt"

func main() {
	var i, j int = 0, 0
	fmt.Println("Give 2 numbers")
	fmt.Scanln(&i)
	fmt.Scanln(&j)
	if i%4 == 0 {
		fmt.Println(i, " is divisible by 4")
	} else {
		fmt.Println(i, " is not divisible by 4")
	}
	if j%4 == 0 {
		fmt.Println(j, " is divisible by 4")
	} else {
		fmt.Println(j, " is not divisible by 4")
	}

	if i%2 == 0 && j%2 == 0 {
		fmt.Println(i, " and ", j, " are even")
	} else if i%2 == 0 {
		fmt.Println(i, " is even and ", j, " is odd")
	} else if j%2 == 0 {
		fmt.Println(j, " is even and ", i, " is odd")
	} else {
		fmt.Println(i, " and ", j, " are odd")
	}

	var num1 int = 0
	fmt.Println("Give another number: ")
	fmt.Scanln(&num1)
	if num1 < 0 {
		fmt.Println(num1, " is negative")
	} else if num1 < 10 {
		fmt.Println(num1, " has 1 digit")
	} else {
		fmt.Println(num1, " has multiple digits")
	}
	var num2 int
	fmt.Println("Give another number: ")
	fmt.Scanln(&num2)
	if num2 < 0 {
		fmt.Println(num2, " is negative")
	} else if num2 < 10 {
		fmt.Println(num2, " has 1 digit")
	} else {
		fmt.Println(num2, " has multiple digits")
	}

	if num1 == num2 {
		fmt.Println(num1, " is equal to ", num2)
	} else {
		fmt.Println(num1, " is unequal to ", num2)
		if i < j {
			fmt.Println(num1, " is less than ", num2)
		} else {
			fmt.Println(num1, " is more than ", num2)
		}
	}
}
