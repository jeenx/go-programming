/*
Hands-on exercise #9
Create a program that uses a switch statement with the switch expression
specified as a variable of TYPE string with the IDENTIFIER “favSport”.
*/
package main

import "fmt"

func main() {
	favSport := "football"
	switch favSport {
	case "soccer":
		fmt.Println("Lets go watch some soccer")
	case "football":
		fmt.Printf("Lets go watch some football")
	}

}
