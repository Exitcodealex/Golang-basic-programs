package main

import "fmt"

func main() {
	i := 0
	for i <= 5 {
		fmt.Println("i=", i)
		i = i + 1
	}

	for j := 0; j < 5; j = j + 2 {
		fmt.Println("j=", j)
		j--
		fmt.Println("j=", j)
	}
	for i := range 6 {
		fmt.Println("range", i)
	}
	for {
		fmt.Println("infiniteloop")
		break
	}
	for n := range 6 {
		if n%3 == 0 {

			continue
		}
		fmt.Println(n)
	}
}
