package main

import "fmt"

func main() {
	var firstName, secondName string
	var age int
	fmt.Scan(&firstName, &secondName, &age)
	fmt.Println("Имя: ", firstName, ", Фамилия: ", secondName, ", Возраст: ", age, ". Студент BPS")
}
