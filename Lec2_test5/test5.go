package main

import (
	"fmt"
)

func main() {

	var a, b, c string

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	if (a == "раз" || a == "один" || a == "1") && (b == "два" || b == "2") && (c == "три" || c == "3") {
		fmt.Println("ОК")
	} else {
		fmt.Println("НЕ ПРАВИЛЬНО")
	}
}
