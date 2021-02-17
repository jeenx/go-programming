/*
Hands-on exercise #7
Building on the previous hands-on exercise, create a program that uses
“else if” and “else”.
*/
package main

import "fmt"

func main() {
	x := 3

	if x == 2 {
		fmt.Println("x is equal to 2")
	} else if x == 3 {
		fmt.Println("x is equal to 3")
	} else {
		fmt.Println("x has not been initialized")
	}
}
