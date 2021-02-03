/*
Hands-on exercise #5
Create a variable of type string using a raw string literal. Print it.
*/
package main

import "fmt"

var rawString string = `This is a raw $tring`

func main() {
	fmt.Println(rawString)
}
