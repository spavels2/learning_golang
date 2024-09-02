package main

import "fmt"

func main() {

	var a, b float64

	fmt.Scan(&a)
	fmt.Scan(&b)

	var sum = int(a + b)

	if sum%2 == 0 {
		fmt.Println("Сумма двух чисел четно")
	} else {
		fmt.Println("Сумма двух чисел не четно")
	}
}
