/*
Hands-on exercise #3
Create TYPED and UNTYPED constants. Print the values of the constants.
*/
package main

import "fmt"

const (
	a     = 42
	b int = 42
)

func main() {
	fmt.Println(a, b)
}
