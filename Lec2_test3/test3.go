package main

import (
	"fmt"
	"strings"
)

func main() {

	var str, word string

	word = "чат"
	fmt.Scan(&str)

	if strings.Contains(str, word) {
		fmt.Println("БОТ")
	} else {
		fmt.Println("НЕ БОТ")
	}
}
