/*
Hands-on exercise #3
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.
*/
package main

import "fmt"

func main() {
	birthYear := 1991

	for birthYear <= 2021 {
		fmt.Println(birthYear)
		birthYear++
	}
}
