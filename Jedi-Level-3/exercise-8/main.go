/*
Hands-on exercise #8
Create a program that uses a switch statement with no switch expression specified.
*/
package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("True case")
	case false:
		fmt.Println("False case")
	default:
		fmt.Println("Default case")
	}
}
