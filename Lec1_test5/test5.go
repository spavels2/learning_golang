package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a)
	fmt.Scan(&b)

	c = a*2 + b*2

	fmt.Println(c * 2)
}
