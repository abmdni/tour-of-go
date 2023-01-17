// Constants are declared like variables, but with the const keyword.

// Constants can be character, string, boolean, or numeric values.

// Constants cannot be declared using the := syntax.

package main

import "fmt"

const Pi = 3.14

func main() {
	const stringg = "abc"

	fmt.Println("helllo strings ", stringg)
	fmt.Println("happy", Pi, "day")
	// stringg = "nanana" not working
	fmt.Println("nah")
}
