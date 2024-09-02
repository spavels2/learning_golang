package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	// var aa = fmt.Sprintf("%d", a)
	// var bb = fmt.Sprintf("%d", b)
	// var cc = fmt.Sprintf("%d", c)

	// fmt.Println(cc + bb + aa)

	fmt.Println(c*100 + b*10 + a)
}
