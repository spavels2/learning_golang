package main

import "fmt"

func main() {

	var word1, word2, word3, word4, autor string

	fmt.Scan(&word1, &word2, &word3, &word4, &autor)

	fmt.Println(word4, " - ", autor)
	fmt.Println(word3, " - ", autor)
	fmt.Println(word2, " - ", autor)
	fmt.Println(word1, " - ", autor)
}
