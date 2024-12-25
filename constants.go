package main

import (
	"fmt"
	"math"
)

var s string = "constant"

func main() {
	fmt.Println(s)
	const n = 5e41
	const d = 3e40 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
