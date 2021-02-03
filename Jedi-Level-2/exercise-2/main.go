/*
Using the following operators, write expressions and assign their values to
variables:
	g. ==
	h. <=
	i. >=
	j. !=
	k. <
	l. >
Now print each of the variables.
*/
package main

import "fmt"

func main() {
	g := (42 == 42)
	h := (42 <= 42)
	i := (42 >= 42)
	j := (42 != 42)
	k := (42 < 42)
	l := (42 > 42)

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}
