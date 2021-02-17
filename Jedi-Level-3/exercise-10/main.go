/*
Hands-on exercise #10
Write down what these print:
	1. fmt.Println(true && true)
	2. fmt.Println(true && false)
	3. fmt.Println(true || true)
	4. fmt.Println(true || false)
	5. fmt.Println(!true)
*/
package main

import "fmt"

func main() {
	fmt.Println(true && true)  // true
	fmt.Println(true && false) // false
	fmt.Println(true || true)  // true
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false
}
