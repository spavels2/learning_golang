package main

import "fmt"

func main() {
	var drink, meal, subMeal, time string
	fmt.Scan(&drink, &meal, &subMeal, &time)
	fmt.Print("I wanna some ", drink)
	fmt.Print(", my favorite meal - ", meal)
	fmt.Print(". Give me also ", subMeal)
	fmt.Print(". I will come soon... ", time)
}
