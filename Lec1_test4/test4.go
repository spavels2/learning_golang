package main

import "fmt"

func main() {

	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println("Периметр: ", a+a+b+b)
	fmt.Println("Площадь: ", a*b)
}
