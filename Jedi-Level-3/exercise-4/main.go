/*
Hands-on exercise #4
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.
*/
package main

import "fmt"

func main() {
	birthYear := 1991

	for {
		if birthYear > 2021 {
			break
		}

		fmt.Println(birthYear)

		birthYear++
	}
}
